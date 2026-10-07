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

type stringSet map[string]struct{}

func (set stringSet) contains(value string) bool {
	_, ok := set[value]
	return ok
}

var datePaths = stringSet{
	"localTime": {},
	"logicalSessionRecordCache.lastSessionsCollectionJobTimestamp": {},
	"logicalSessionRecordCache.lastTransactionReaperJobTimestamp":  {},
	"repl.lastWrite.lastWriteDate":                                 {},
	"repl.lastWrite.majorityWriteDate":                             {},
	"security.SSLServerCertificateExpirationDate":                  {},
}
var objectIDPaths = stringSet{"repl.topologyVersion.processId": {}, "repl.electionId": {}, "$gleStats.electionId": {}}
var timestampPaths = stringSet{
	"$clusterTime.clusterTime":                                 {},
	"$configServerState.opTime.ts":                             {},
	"$gleStats.lastOpTime":                                     {},
	"lastCommittedOpTime":                                      {},
	"operationTime":                                            {},
	"repl.lastWrite.majorityOpTime.ts":                         {},
	"repl.lastWrite.opTime.ts":                                 {},
	"sharding.lastSeenConfigServerOpTime.ts":                   {},
	"storageEngine.oldestRequiredTimestampForCrashRecovery":    {},
	"wiredTiger.oplog.visibility timestamp":                    {},
	"wiredTiger.snapshot-window-settings.min pinned timestamp": {},
}
var binaryPaths = stringSet{"$clusterTime.signature.hash": {}}
var forbiddenIdentityFragments = []string{"evergage.com", "salesforce.com", "prod4", "p4m1ssd", "rs_prod"}

const (
	syntheticDate             = "2020-01-02T03:04:05.000Z"
	syntheticSessionDate      = "2020-01-02T03:03:05.000Z"
	syntheticCertificateDate  = "2020-02-01T00:00:00.000Z"
	syntheticSnapshotLatest   = "Jan 02 03:04:05:7"
	syntheticSnapshotOldest   = "Jan 02 02:59:05:7"
	syntheticObjectID         = "000000000000000000000000"
	syntheticBinary           = "AAAAAAAAAAAAAAAAAAAAAAAAAAA="
	syntheticHost             = "fixture-host"
	syntheticMember           = "fixture-member-1:27017"
	syntheticReplicaSet       = "fixture-replica-set"
	syntheticConfigConnection = "fixture-config/fixture-config-1:27017,fixture-config-2:27017,fixture-config-3:27017"
	syntheticCertificate      = "CN=fixture"
)

// SanitizeJSON replaces captured runtime state with deterministic fixture values.
// It preserves every map, array, scalar path, and source scalar type.
func SanitizeJSON(source []byte) ([]byte, error) {
	value, err := decodeJSON(source)
	if err != nil {
		return nil, err
	}
	return marshalIndent(sanitize(value, ""))
}

// ValidateSanitizedJSON rejects source fixtures that do not contain the
// deterministic numeric and identity/runtime values produced by SanitizeJSON.
func ValidateSanitizedJSON(source []byte) error {
	value, err := decodeJSON(source)
	if err != nil {
		return err
	}
	return validateSanitized(value, "")
}

// ReconstructJSON converts a sanitized fixture into deterministic Canonical Extended JSON.
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

func sanitize(value interface{}, path string) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(v))
		for key, child := range v {
			out[key] = sanitize(child, join(path, key))
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(v))
		for i, child := range v {
			out[i] = sanitize(child, fmt.Sprintf("%s[%d]", path, i))
		}
		return out
	case json.Number:
		return sanitizeNumber(path, v)
	case string:
		return sanitizeString(path, v)
	default:
		return value
	}
}

func sanitizeNumber(path string, value json.Number) json.Number {
	text := value.String()
	if path == "$clusterTime.clusterTime[0]" {
		return exactNumberOfSameJSONType(text, 1577934245)
	}
	if path == "$clusterTime.clusterTime[1]" {
		return exactNumberOfSameJSONType(text, 7)
	}
	if approvedZeroPath(path) && numberIsZero(value) {
		return zeroOfSameJSONType(text)
	}
	if path == "ok" {
		return numberOfSameJSONType(text, 1)
	}

	synthetic := int64(2)
	for index, character := range []byte(path) {
		synthetic += int64(index+1) * int64(character)
	}
	synthetic = synthetic%89 + 2
	if strings.HasPrefix(text, "-") {
		synthetic = -synthetic
	}
	return numberOfSameJSONType(text, synthetic)
}

func approvedZeroPath(path string) bool {
	return strings.HasPrefix(path, "metrics.commands.")
}

func numberIsZero(value json.Number) bool {
	number, err := strconv.ParseFloat(value.String(), 64)
	return err == nil && number == 0
}

func zeroOfSameJSONType(text string) json.Number {
	if strings.ContainsAny(text, ".eE") {
		return json.Number("0.0")
	}
	return json.Number("0")
}

func numberOfSameJSONType(text string, value int64) json.Number {
	if strings.ContainsAny(text, ".eE") {
		return json.Number(strconv.FormatFloat(float64(value)+0.5, 'f', 1, 64))
	}
	return json.Number(strconv.FormatInt(value, 10))
}

func exactNumberOfSameJSONType(text string, value int64) json.Number {
	integer := strconv.FormatInt(value, 10)
	if strings.ContainsAny(text, ".eE") {
		return json.Number(integer + ".0")
	}
	return json.Number(integer)
}

func sanitizeString(path, value string) string {
	basePath := path
	if index := strings.LastIndexByte(basePath, '['); index >= 0 {
		basePath = basePath[:index]
	}
	if binaryPaths.contains(basePath) {
		return syntheticBinary
	}
	if objectIDPaths.contains(basePath) {
		return syntheticObjectID
	}
	switch basePath {
	case "localTime", "repl.lastWrite.lastWriteDate", "repl.lastWrite.majorityWriteDate":
		return syntheticDate
	case "logicalSessionRecordCache.lastSessionsCollectionJobTimestamp", "logicalSessionRecordCache.lastTransactionReaperJobTimestamp":
		return syntheticSessionDate
	case "security.SSLServerCertificateExpirationDate":
		return syntheticCertificateDate
	case "host":
		return syntheticHost
	case "repl.hosts":
		return syntheticMember
	case "repl.me", "repl.primary":
		return syntheticMember
	case "repl.setName":
		return syntheticReplicaSet
	case "security.SSLServerSubjectName":
		return syntheticCertificate
	case "sharding.configsvrConnectionString":
		return syntheticConfigConnection
	case "wiredTiger.snapshot-window-settings.latest majority snapshot timestamp available":
		return syntheticSnapshotLatest
	case "wiredTiger.snapshot-window-settings.oldest majority snapshot timestamp available":
		return syntheticSnapshotOldest
	case "metrics.repl.stateTransition.lastStateTransition":
		if value == "" {
			return ""
		}
		return "stepUp"
	case "repl.primaryOnlyServices.RenameCollectionParticipantService.state",
		"repl.primaryOnlyServices.ReshardingDonorService.state",
		"repl.primaryOnlyServices.ReshardingRecipientService.state",
		"repl.primaryOnlyServices.ShardingDDLCoordinator.state":
		if value == "running" {
			return "running"
		}
		return "paused"
	case "oplogTruncation.processingMethod":
		return "sampling"
	case "process":
		return "mongod"
	case "storageEngine.name":
		return "wiredTiger"
	case "version":
		return "5.0.34"
	case "wiredTiger.uri":
		return "statistics:"
	default:
		if strings.HasPrefix(basePath, "metrics.apiVersions") {
			return "default"
		}
		return "synthetic"
	}
}

func validateSanitized(value interface{}, path string) error {
	switch v := value.(type) {
	case map[string]interface{}:
		for key, child := range v {
			if err := validateSanitized(child, join(path, key)); err != nil {
				return err
			}
		}
	case []interface{}:
		for i, child := range v {
			if err := validateSanitized(child, fmt.Sprintf("%s[%d]", path, i)); err != nil {
				return err
			}
		}
	case json.Number:
		if got := sanitizeNumber(path, v); got.String() != v.String() {
			return fmt.Errorf("%s has non-synthetic numeric value %s; want %s", path, v, got)
		}
	case string:
		if want := sanitizeString(path, v); v != want {
			return fmt.Errorf("%s has non-synthetic string value %q", path, v)
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
	if timestampPaths.contains(path) {
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
		if datePaths.contains(path) {
			if parsed, err := time.Parse(time.RFC3339Nano, v); err == nil {
				return map[string]interface{}{"$date": map[string]interface{}{"$numberLong": strconv.FormatInt(parsed.UnixNano()/int64(time.Millisecond), 10)}}
			}
		}
		if objectIDPaths.contains(path) && len(v) == 24 {
			return map[string]interface{}{"$oid": v}
		}
		if binaryPaths.contains(path) {
			if _, err := base64.StdEncoding.DecodeString(v); err == nil {
				return map[string]interface{}{"$binary": map[string]interface{}{"base64": v, "subType": "00"}}
			}
		}
		return v
	default:
		return value
	}
}
func reconstructTimestamp(path string, value interface{}) (map[string]interface{}, bool) {
	switch v := value.(type) {
	case []interface{}:
		if len(v) != 2 {
			return nil, false
		}
		return map[string]interface{}{"t": v[0], "i": v[1]}, true
	case string:
		if path != "wiredTiger.snapshot-window-settings.latest majority snapshot timestamp available" &&
			path != "wiredTiger.snapshot-window-settings.oldest majority snapshot timestamp available" {
			return nil, false
		}
		separator := strings.LastIndexByte(v, ':')
		if separator < 0 {
			return nil, false
		}
		seconds, err := time.Parse("2006 Jan 02 15:04:05", "2020 "+v[:separator])
		if err != nil {
			return nil, false
		}
		increment, err := strconv.ParseInt(v[separator+1:], 10, 64)
		if err != nil {
			return nil, false
		}
		return map[string]interface{}{"t": seconds.Unix(), "i": increment}, true
	default:
		return nil, false
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
