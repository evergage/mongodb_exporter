package fixturetool

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	RedactionVersion      = "1"
	ReconstructionVersion = "1"
	NormalizationVersion  = "1"
)

type Fixture struct {
	Role                string `json:"role"`
	MongoDBVersion      string `json:"mongodb_version"`
	CaptureTime         string `json:"capture_time"`
	Source              string `json:"source"`
	SourceSHA256        string `json:"source_sha256"`
	Reconstructed       string `json:"reconstructed"`
	ReconstructedSHA256 string `json:"reconstructed_sha256"`
	Normalized          string `json:"normalized"`
	NormalizedSHA256    string `json:"normalized_sha256"`
}

type Manifest struct {
	RedactionVersion           string    `json:"redaction_version"`
	ReconstructionVersion      string    `json:"reconstruction_version"`
	NormalizationVersion       string    `json:"normalization_version"`
	OriginalBSONTypesAvailable bool      `json:"original_bson_types_available"`
	Fixtures                   []Fixture `json:"fixtures"`
}

type fixtureSpec struct{ role, file string }

var specs = []fixtureSpec{{"primary", "server-status-primary-5.0.34.json"}, {"secondary", "server-status-secondary-5.0.34.json"}}
var datePaths = map[string]bool{
	"localTime": true,
	"logicalSessionRecordCache.lastSessionsCollectionJobTimestamp": true,
	"logicalSessionRecordCache.lastTransactionReaperJobTimestamp":  true,
	"repl.lastWrite.lastWriteDate":                                 true,
	"repl.lastWrite.majorityWriteDate":                             true,
	"security.SSLServerCertificateExpirationDate":                  true,
}
var objectIDPaths = map[string]bool{
	"repl.topologyVersion.processId": true,
	"repl.electionId":                true,
	"$gleStats.electionId":           true,
}
var timestampPaths = map[string]bool{
	"$clusterTime.clusterTime": true,
	"$gleStats.lastOpTime":     true,
}
var binaryPaths = map[string]bool{"$clusterTime.signature.hash": true}
var forbiddenIdentityFragments = []string{"evergage.com", "salesforce.com", "prod4", "p4m1ssd", "rs_prod"}

func Generate(root string, check bool) error {
	manifest := Manifest{RedactionVersion: RedactionVersion, ReconstructionVersion: ReconstructionVersion, NormalizationVersion: NormalizationVersion, OriginalBSONTypesAvailable: false}
	for _, spec := range specs {
		sourceRel := filepath.Join("source", spec.file)
		sourcePath := filepath.Join(root, sourceRel)
		source, err := ioutil.ReadFile(sourcePath)
		if err != nil {
			return err
		}
		if err := rejectIdentities(source); err != nil {
			return fmt.Errorf("%s: %w", sourceRel, err)
		}
		value, err := decodeJSON(source)
		if err != nil {
			return fmt.Errorf("%s: %w", sourceRel, err)
		}
		doc, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%s: root is not an object", sourceRel)
		}
		version, _ := doc["version"].(string)
		captureTime, _ := doc["localTime"].(string)
		reconstructed := reconstruct(value, "")
		extJSON, err := marshalIndent(reconstructed)
		if err != nil {
			return err
		}
		normalized, err := normalize(reconstructed, "")
		if err != nil {
			return err
		}
		normalizedJSON, err := marshalIndent(normalized)
		if err != nil {
			return err
		}
		if !jsonEqual(source, normalizedJSON) {
			return fmt.Errorf("%s: reconstructed normalization changed source values", sourceRel)
		}
		reconstructedRel := filepath.Join("reconstructed", spec.file)
		normalizedRel := filepath.Join("normalized", spec.file)
		if err := emit(filepath.Join(root, reconstructedRel), extJSON, check); err != nil {
			return err
		}
		if err := emit(filepath.Join(root, normalizedRel), normalizedJSON, check); err != nil {
			return err
		}
		manifest.Fixtures = append(manifest.Fixtures, Fixture{
			Role: spec.role, MongoDBVersion: version, CaptureTime: captureTime,
			Source: filepath.ToSlash(sourceRel), SourceSHA256: digest(source),
			Reconstructed: filepath.ToSlash(reconstructedRel), ReconstructedSHA256: digest(extJSON),
			Normalized: filepath.ToSlash(normalizedRel), NormalizedSHA256: digest(normalizedJSON),
		})
	}
	manifestJSON, err := marshalIndent(manifest)
	if err != nil {
		return err
	}
	return emit(filepath.Join(root, "manifest.json"), manifestJSON, check)
}

func ReconstructJSON(source []byte) ([]byte, error) {
	value, err := decodeJSON(source)
	if err != nil {
		return nil, err
	}
	return marshalIndent(reconstruct(value, ""))
}

func rejectIdentities(data []byte) error {
	lower := strings.ToLower(string(data))
	for _, fragment := range forbiddenIdentityFragments {
		if strings.Contains(lower, strings.ToLower(fragment)) {
			return fmt.Errorf("contains forbidden identity fragment %q", fragment)
		}
	}
	return nil
}

func decodeJSON(data []byte) (interface{}, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	var value interface{}
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	return value, nil
}

func reconstruct(value interface{}, path string) interface{} {
	if timestampPaths[path] {
		values, ok := value.([]interface{})
		if ok && len(values) == 2 {
			return map[string]interface{}{"$timestamp": map[string]interface{}{"t": values[0], "i": values[1]}}
		}
	}
	switch v := value.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(v))
		for key, child := range v {
			out[key] = reconstruct(child, join(path, key))
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(v))
		for i, child := range v {
			out[i] = reconstruct(child, path)
		}
		return out
	case json.Number:
		text := v.String()
		if strings.ContainsAny(text, ".eE") {
			return map[string]interface{}{"$numberDouble": text}
		}
		return map[string]interface{}{"$numberLong": text}
	case string:
		if datePaths[path] {
			parsed, err := time.Parse(time.RFC3339Nano, v)
			if err == nil {
				return map[string]interface{}{"$date": map[string]interface{}{"$numberLong": strconv.FormatInt(parsed.UnixNano()/int64(time.Millisecond), 10)}}
			}
		}
		if objectIDPaths[path] && len(v) == 24 {
			return map[string]interface{}{"$oid": v}
		}
		if binaryPaths[path] {
			if _, err := base64.StdEncoding.DecodeString(v); err == nil {
				return map[string]interface{}{"$binary": map[string]interface{}{"base64": v, "subType": "00"}}
			}
		}
		return v
	default:
		return value
	}
}

func normalize(value interface{}, path string) (interface{}, error) {
	switch v := value.(type) {
	case map[string]interface{}:
		if raw, ok := v["$numberLong"]; ok && len(v) == 1 {
			return json.Number(raw.(string)), nil
		}
		if raw, ok := v["$numberDouble"]; ok && len(v) == 1 {
			return json.Number(raw.(string)), nil
		}
		if raw, ok := v["$oid"]; ok && len(v) == 1 {
			return raw.(string), nil
		}
		if raw, ok := v["$date"]; ok && len(v) == 1 {
			inner := raw.(map[string]interface{})
			millis, err := strconv.ParseInt(inner["$numberLong"].(string), 10, 64)
			if err != nil {
				return nil, err
			}
			return time.Unix(0, millis*int64(time.Millisecond)).UTC().Format(time.RFC3339Nano), nil
		}
		if raw, ok := v["$timestamp"]; ok && len(v) == 1 {
			inner := raw.(map[string]interface{})
			return []interface{}{inner["t"], inner["i"]}, nil
		}
		if raw, ok := v["$binary"]; ok && len(v) == 1 {
			inner := raw.(map[string]interface{})
			return inner["base64"], nil
		}
		out := make(map[string]interface{}, len(v))
		for key, child := range v {
			n, err := normalize(child, join(path, key))
			if err != nil {
				return nil, err
			}
			out[key] = n
		}
		return out, nil
	case []interface{}:
		out := make([]interface{}, len(v))
		for i, child := range v {
			n, err := normalize(child, path)
			if err != nil {
				return nil, err
			}
			out[i] = n
		}
		return out, nil
	default:
		return value, nil
	}
}

func join(path, key string) string {
	if path == "" {
		return key
	}
	return path + "." + key
}
func marshalIndent(value interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}
func digest(data []byte) string { sum := sha256.Sum256(data); return hex.EncodeToString(sum[:]) }
func jsonEqual(a, b []byte) bool {
	av, e1 := decodeJSON(a)
	bv, e2 := decodeJSON(b)
	return e1 == nil && e2 == nil && fmt.Sprintf("%#v", av) == fmt.Sprintf("%#v", bv)
}
func emit(path string, data []byte, check bool) error {
	if check {
		current, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		if !bytes.Equal(current, data) {
			return fmt.Errorf("%s is stale; run serverstatus-fixtures", path)
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func SortedLeafPaths(value interface{}) []string {
	var paths []string
	walkLeaves(value, "", &paths)
	sort.Strings(paths)
	return paths
}
func walkLeaves(value interface{}, path string, paths *[]string) {
	switch v := value.(type) {
	case map[string]interface{}:
		if len(v) == 0 {
			*paths = append(*paths, path)
			return
		}
		for k, c := range v {
			walkLeaves(c, join(path, k), paths)
		}
	case []interface{}:
		if len(v) == 0 {
			*paths = append(*paths, path)
			return
		}
		for i, c := range v {
			walkLeaves(c, fmt.Sprintf("%s[%d]", path, i), paths)
		}
	default:
		*paths = append(*paths, path)
	}
}
