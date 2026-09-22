package fixturetool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var datePaths = map[string]bool{
	"localTime": true,
	"logicalSessionRecordCache.lastSessionsCollectionJobTimestamp": true,
	"logicalSessionRecordCache.lastTransactionReaperJobTimestamp":  true,
	"repl.lastWrite.lastWriteDate":                                 true,
	"repl.lastWrite.majorityWriteDate":                             true,
	"security.SSLServerCertificateExpirationDate":                  true,
}
var objectIDPaths = map[string]bool{"repl.topologyVersion.processId": true, "repl.electionId": true, "$gleStats.electionId": true}
var timestampPaths = map[string]bool{"$clusterTime.clusterTime": true, "$gleStats.lastOpTime": true}
var binaryPaths = map[string]bool{"$clusterTime.signature.hash": true}
var forbiddenIdentityFragments = []string{"evergage.com", "salesforce.com", "prod4", "p4m1ssd", "rs_prod"}

// ReconstructJSON converts the sanitized fixture into deterministic Canonical Extended JSON.
func ReconstructJSON(source []byte) ([]byte, error) {
	if err := rejectIdentities(source); err != nil {
		return nil, err
	}
	value, err := decodeJSON(source)
	if err != nil {
		return nil, err
	}
	return marshalIndent(reconstruct(value, ""))
}

func rejectIdentities(data []byte) error {
	lower := strings.ToLower(string(data))
	for _, fragment := range forbiddenIdentityFragments {
		if strings.Contains(lower, fragment) {
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
		if values, ok := value.([]interface{}); ok && len(values) == 2 {
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
			if parsed, err := time.Parse(time.RFC3339Nano, v); err == nil {
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

type Leaf struct {
	Path     string
	Segments []string
	Type     string
}

func SortedLeaves(value interface{}) []Leaf {
	var leaves []Leaf
	walkTypedLeaves(value, "", nil, &leaves)
	sort.Slice(leaves, func(i, j int) bool { return leaves[i].Path < leaves[j].Path })
	return leaves
}

func SortedLeafPaths(value interface{}) []string {
	leaves := SortedLeaves(value)
	paths := make([]string, len(leaves))
	for i := range leaves {
		paths[i] = leaves[i].Path
	}
	return paths
}

func sourceKeyPath(path, key string) string {
	if key != "" {
		valid := true
		for i, r := range key {
			if !(r == '_' || r == '$' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || i > 0 && r >= '0' && r <= '9') {
				valid = false
				break
			}
		}
		if valid {
			if path == "" {
				return key
			}
			return path + "." + key
		}
	}
	escaped := strings.ReplaceAll(strings.ReplaceAll(key, "\\", "\\\\"), "'", "\\'")
	return path + "['" + escaped + "']"
}

func sourceType(value interface{}) string {
	switch value.(type) {
	case json.Number:
		return "number"
	case string:
		return "string"
	case bool:
		return "boolean"
	case nil:
		return "null"
	default:
		return "container"
	}
}

func walkTypedLeaves(value interface{}, path string, segments []string, leaves *[]Leaf) {
	switch v := value.(type) {
	case map[string]interface{}:
		if len(v) == 0 {
			*leaves = append(*leaves, Leaf{path, append([]string(nil), segments...), "object"})
			return
		}
		for key, child := range v {
			next := append(append([]string(nil), segments...), key)
			walkTypedLeaves(child, sourceKeyPath(path, key), next, leaves)
		}
	case []interface{}:
		if len(v) == 0 {
			*leaves = append(*leaves, Leaf{path, append([]string(nil), segments...), "array"})
			return
		}
		for i, child := range v {
			index := strconv.Itoa(i)
			next := append(append([]string(nil), segments...), index)
			walkTypedLeaves(child, fmt.Sprintf("%s[%d]", path, i), next, leaves)
		}
	default:
		*leaves = append(*leaves, Leaf{path, append([]string(nil), segments...), sourceType(value)})
	}
}
