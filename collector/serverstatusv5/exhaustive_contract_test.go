package serverstatusv5

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFixtureLeafInventoryPreservesLiteralKeysAndTypes(t *testing.T) {
	contract, paths := loadCoverageInputs(t)
	if len(paths) != 1581 || len(contract.Entries) != 1581 {
		t.Fatalf("paths=%d entries=%d, want 1581 each", len(paths), len(contract.Entries))
	}
	seen := map[string]bool{}
	for _, leaf := range paths {
		seen[leaf.Path] = true
	}
	for _, path := range []string{"metrics.apiVersions[''][0]", "transportSecurity['1.2']"} {
		if !seen[path] {
			t.Errorf("literal-key path %q missing", path)
		}
	}
	data, err := ioutil.ReadFile("testdata/source/server-status-primary-5.0.34.json")
	if err != nil {
		t.Fatal(err)
	}
	var value interface{}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.UseNumber()
	if err := decoder.Decode(&value); err != nil {
		t.Fatal(err)
	}
	leaves := fixturetool.SortedLeaves(value)
	types := map[string]string{}
	for _, leaf := range leaves {
		types[leaf.Path] = leaf.Type
	}
	if types["ok"] != "number" || types["host"] != "string" || types["trafficRecording.running"] != "boolean" {
		t.Fatalf("unexpected source types: ok=%q host=%q running=%q", types["ok"], types["host"], types["trafficRecording.running"])
	}
}

func TestGeneratedNumericPolicyAndExceptions(t *testing.T) {
	contract := loadPolicyContract(t)
	counts := map[string]int{}
	byPath := map[string]ContractEntry{}
	for _, entry := range contract.Entries {
		byPath[entry.Path] = entry
		if strings.HasPrefix(entry.DecisionSource, "numeric-mapping-proposals.md") {
			counts[entry.Metric.Type]++
		}
		if entry.Metric != nil && entry.DecisionSource != "existing modern contract" {
			if entry.Metric.Conversion != 1 {
				t.Errorf("%s conversion=%v", entry.Path, entry.Metric.Conversion)
			}
			if entry.Metric.Type == "counter" && !strings.HasSuffix(entry.Metric.Family, "_total") {
				t.Errorf("counter %s lacks suffix", entry.Metric.Family)
			}
		}
	}
	if counts["counter"] != 411 || counts["gauge"] != 248 {
		t.Fatalf("review-corrected generated types=%v", counts)
	}
	for _, entry := range contract.Entries {
		if !strings.Contains(entry.DecisionSource, "BSON Timestamp state") {
			continue
		}
		if entry.Metric == nil || entry.Metric.Type != "gauge" || entry.Metric.Transform == "" || strings.HasSuffix(entry.Metric.Family, "_total") {
			t.Errorf("logical timestamp %s has invalid mapping %+v", entry.Path, entry.Metric)
		}
	}
	checks := map[string]string{
		"connections.exhaustHello":                                         "gauge",
		"transportSecurity['1.2']":                                         "counter",
		"wiredTiger['data-handle']['connection data handle size']":         "gauge",
		"wiredTiger.transaction['transaction checkpoint generation']":      "gauge",
		"wiredTiger.transaction['prepared transactions']":                  "counter",
		"wiredTiger.session['tiered storage local retention time (secs)']": "gauge",
		"wiredTiger.session['table compact timeout']":                      "counter",
		"wiredTiger.connection['hash bucket array size general']":          "gauge",
	}
	for path, want := range checks {
		entry := byPath[path]
		if entry.Metric == nil || entry.Metric.Type != want {
			t.Errorf("%s type=%v want %s", path, entry.Metric, want)
		}
	}
}
func TestGeneratedDefinitionsUsePreSplitPaths(t *testing.T) {
	for _, definition := range metricDefinitions {
		if len(definition.PathSegments) == 0 {
			t.Fatalf("%s has no pre-split path", definition.SourcePath)
		}
	}
	segmentCount := 0
	allocations := testing.AllocsPerRun(100, func() {
		count := 0
		for _, definition := range metricDefinitions {
			count += len(definition.PathSegments)
		}
		segmentCount = count
	})
	if allocations != 0 || segmentCount == 0 {
		t.Fatalf("pre-split path preparation allocated %v objects for %d segments", allocations, segmentCount)
	}
}

func TestDynamicVocabulariesAndUnknownDiagnostics(t *testing.T) {
	contract := loadPolicyContract(t)
	want := map[string]int{"commands": 60, "aggregation_stages": 49, "operators_expressions": 141, "operators_groupAccumulators": 14, "operators_match": 35, "operators_windowAccumulators": 20}
	for name, size := range want {
		if got := len(contract.DynamicVocabularies[name]); got != size {
			t.Errorf("%s=%d want %d", name, got, size)
		}
	}

	data, err := bson.Marshal(bson.M{"metrics": bson.M{
		"commands":         bson.M{"notReviewed": bson.M{"total": 1}, "emptyUnknown": bson.M{}},
		"aggStageCounters": bson.M{"$notReviewed": 1},
		"operatorCounters": bson.M{"notReviewedCategory": bson.M{"$eq": 1}, "expressions": bson.M{"$notReviewed": 1}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	collector := &rawModuleCollector{module: New(), raw: bson.Raw(data)}
	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(collector); err != nil {
		t.Fatal(err)
	}
	families, err := registry.Gather()
	if err != nil {
		t.Fatal(err)
	}
	values := map[string]float64{}
	for _, family := range families {
		if family.GetName() != "mongodb_server_status_unknown_keys_total" {
			continue
		}
		for _, metric := range family.Metric {
			if len(metric.Label) == 1 {
				values[metric.Label[0].GetValue()] = metric.Counter.GetValue()
			}
		}
	}
	if values["commands"] != 2 || values["aggregation_stages"] != 1 || values["operators"] != 2 {
		t.Fatalf("unknown diagnostics=%v", values)
	}
}

func TestFixedHistogramMappings(t *testing.T) {
	contract := loadPolicyContract(t)
	histogramLeaves, emittedBuckets := 0, 0
	for _, entry := range contract.Entries {
		if entry.Reason != "Observed fixed histogram range." {
			continue
		}
		histogramLeaves++
		if entry.Metric == nil {
			continue
		}
		lowerBounds := 0
		for name, value := range entry.Metric.Labels {
			if strings.HasPrefix(name, "lower_bound_") {
				lowerBounds++
				if _, err := strconv.ParseFloat(value, 64); err != nil {
					t.Errorf("%s %s=%q", entry.Path, name, value)
				}
			}
		}
		if !strings.HasSuffix(entry.Path, ".totalCount") {
			emittedBuckets++
			if lowerBounds != 1 {
				t.Errorf("%s lower-bound labels=%d", entry.Path, lowerBounds)
			}
		}
	}
	if histogramLeaves != 151 || emittedBuckets != 93 {
		t.Fatalf("histogram leaves=%d emitted buckets=%d", histogramLeaves, emittedBuckets)
	}
	for description, want := range map[string]string{
		"eviction walk target pages histogram - 128 and higher":   "128",
		"file system read latency histogram (bucket 2) - 50-99ms": "50",
		"operation write latency histogram (bucket 5) - 10000us+": "10000",
	} {
		if got, err := wiredTigerHistogramLowerBound(description); err != nil || got != want {
			t.Errorf("bound(%q)=%q,%v want %q", description, got, err, want)
		}
	}
	if _, err := wiredTigerHistogramLowerBound("histogram without a range"); err == nil {
		t.Fatal("ambiguous WiredTiger range accepted")
	}
}

func TestReviewedStringsAndDuplicateOwnership(t *testing.T) {
	contract := loadPolicyContract(t)
	drops := 0
	modern := map[string]bool{}
	for _, entry := range contract.Entries {
		if entry.Classification == "drop" {
			drops++
			if !entry.Reviewed || entry.Reason == "" {
				t.Errorf("unreviewed drop %s", entry.Path)
			}
		}
		if entry.Metric != nil {
			modern[entry.Path] = true
		}
	}
	if drops != 28 {
		t.Fatalf("reviewed drops=%d want 28", drops)
	}
	for _, path := range []string{"$clusterTime.signature.keyId", "pid", "metrics.cursor.open.noTimeout", "metrics.cursor.open.pinned", "metrics.cursor.open.total", "metrics.cursor.timedOut", "wiredTiger.concurrentTransactions.read.available", "wiredTiger.concurrentTransactions.write.out"} {
		if modern[path] {
			t.Errorf("modern mapping remains for reviewed legacy/drop path %s", path)
		}
	}
}
func TestAnchoredSnapshotTimestampSelection(t *testing.T) {
	parts, err := parseSnapshotParts("Dec 31 23:59:59:42")
	if err != nil {
		t.Fatal(err)
	}
	anchor, ok := anchorLatestSnapshot(parts, time.Date(2027, 1, 1, 0, 0, 1, 0, time.UTC), uint32(time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC).Unix()))
	if !ok || anchor.year != 2026 || anchor.offsetSeconds != 0 {
		t.Fatalf("anchor=%+v ok=%v", anchor, ok)
	}
	value, ok := anchoredSnapshot(parts, anchor, true)
	if !ok || value.increment != 42 {
		t.Fatalf("value=%+v ok=%v", value, ok)
	}
	offsetParts, err := parseSnapshotParts("Sep 21 23:05:28:7")
	if err != nil {
		t.Fatal(err)
	}
	offsetTarget := time.Date(2026, 9, 21, 21, 5, 28, 0, time.UTC)
	offsetAnchor, ok := anchorLatestSnapshot(offsetParts, offsetTarget, uint32(offsetTarget.Unix()))
	if !ok || offsetAnchor.offsetSeconds != 2*60*60 {
		t.Fatalf("timezone anchor=%+v ok=%v", offsetAnchor, ok)
	}
	ambiguousParts, err := parseSnapshotParts("Jan 01 00:00:00:1")
	if err != nil {
		t.Fatal(err)
	}
	ambiguousTarget := time.Date(2026, 1, 1, 0, 7, 30, 0, time.UTC)
	if _, ok := anchorLatestSnapshot(ambiguousParts, ambiguousTarget, uint32(ambiguousTarget.Unix())); ok {
		t.Fatal("equidistant timezone candidates were accepted")
	}
	if _, err := parseSnapshotParts("not a timestamp"); err == nil {
		t.Fatal("malformed timestamp accepted")
	}
	if _, ok := anchorLatestSnapshot(parts, time.Now(), 0); ok {
		t.Fatal("timestamp anchored without logical timestamp")
	}
}

func TestFixtureSnapshotAndStringMetrics(t *testing.T) {
	primary := gatherFixture(t, "primary")
	checks := map[string]float64{
		"mongodb_server_status_repl_last_write_last_write_date_milliseconds":                             1790024728000,
		"mongodb_server_status_repl_last_write_majority_write_date_milliseconds":                         1790024728000,
		"mongodb_server_status_wired_tiger_snapshot_window_latest_majority_snapshot_timestamp_seconds":   1790024728,
		"mongodb_server_status_wired_tiger_snapshot_window_latest_majority_snapshot_timestamp_increment": 2424,
		"mongodb_server_status_wired_tiger_snapshot_window_oldest_majority_snapshot_timestamp_seconds":   1790024428,
		"mongodb_server_status_wired_tiger_snapshot_window_oldest_majority_snapshot_timestamp_increment": 2424,
	}
	for family, want := range checks {
		if got := metricValueDTO(metricWithLabels(primary[family], nil)); got != want {
			t.Errorf("%s=%v want %v", family, got, want)
		}
	}
	method := primary["mongodb_server_status_oplog_truncation_processing_method_info"]
	if method == nil || len(method.Metric) != 1 || len(method.Metric[0].Label) != 1 || method.Metric[0].Label[0].GetName() != "method" {
		t.Fatalf("processing method metric=%v", method)
	}
}

func gatherRaw(t *testing.T, document interface{}) map[string]*dto.MetricFamily {
	t.Helper()
	data, err := bson.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(&rawModuleCollector{module: New(), raw: bson.Raw(data)}); err != nil {
		t.Fatal(err)
	}
	families, err := registry.Gather()
	if err != nil {
		t.Fatal(err)
	}
	result := make(map[string]*dto.MetricFamily, len(families))
	for _, family := range families {
		result[family.GetName()] = family
	}
	return result
}

func TestMethodAndReplicationDatesArePresenceAware(t *testing.T) {
	absent := gatherRaw(t, bson.M{})
	for _, family := range []string{"mongodb_server_status_oplog_truncation_processing_method_info", "mongodb_server_status_repl_last_write_last_write_date_milliseconds", "mongodb_server_status_repl_last_write_majority_write_date_milliseconds"} {
		if absent[family] != nil {
			t.Errorf("absent field emitted %s", family)
		}
	}
	present := gatherRaw(t, bson.M{
		"oplogTruncation": bson.M{"processingMethod": "sampling"},
		"repl":            bson.M{"lastWrite": bson.M{"lastWriteDate": primitive.DateTime(1234), "majorityWriteDate": primitive.DateTime(5678)}},
	})
	method := metricWithLabels(present["mongodb_server_status_oplog_truncation_processing_method_info"], map[string]string{"method": "sampling"})
	if method == nil || metricValueDTO(method) != 1 {
		t.Fatal("reviewed processing method was not emitted")
	}
	unknown := gatherRaw(t, bson.M{"oplogTruncation": bson.M{"processingMethod": "not-reviewed"}})
	if unknown["mongodb_server_status_oplog_truncation_processing_method_info"] != nil {
		t.Fatal("unknown processing method was emitted")
	}
	errorMetric := metricWithLabels(unknown["mongodb_server_status_decode_errors_total"], map[string]string{"family": "mongodb_server_status_oplog_truncation_processing_method_info"})
	if errorMetric == nil || metricValueDTO(errorMetric) != 1 {
		t.Fatal("unknown processing method did not increment bounded diagnostic")
	}
	if got := metricValueDTO(metricWithLabels(present["mongodb_server_status_repl_last_write_last_write_date_milliseconds"], nil)); got != 1234 {
		t.Errorf("lastWriteDate=%v", got)
	}
	if got := metricValueDTO(metricWithLabels(present["mongodb_server_status_repl_last_write_majority_write_date_milliseconds"], nil)); got != 5678 {
		t.Errorf("majorityWriteDate=%v", got)
	}
}
func TestSnapshotAnchoringFailureEmitsZeroSentinels(t *testing.T) {
	families := gatherRaw(t, bson.M{
		"localTime":    time.Unix(1790024728, 0),
		"$clusterTime": bson.M{"clusterTime": primitive.Timestamp{T: 1790024728, I: 1}},
		"wiredTiger": bson.M{"snapshot-window-settings": bson.M{
			"latest majority snapshot timestamp available": "malformed",
			"oldest majority snapshot timestamp available": "also malformed",
		}},
		"asserts": bson.M{"tripwire": 7},
	})
	for _, family := range []string{
		"mongodb_server_status_wired_tiger_snapshot_window_latest_majority_snapshot_timestamp_seconds",
		"mongodb_server_status_wired_tiger_snapshot_window_latest_majority_snapshot_timestamp_increment",
		"mongodb_server_status_wired_tiger_snapshot_window_oldest_majority_snapshot_timestamp_seconds",
		"mongodb_server_status_wired_tiger_snapshot_window_oldest_majority_snapshot_timestamp_increment",
	} {
		if got := metricValueDTO(metricWithLabels(families[family], nil)); got != 0 {
			t.Errorf("%s=%v want zero", family, got)
		}
	}
	if got := metricValueDTO(metricWithLabels(families["mongodb_server_status_asserts_tripwire_total"], nil)); got != 7 {
		t.Errorf("other family suppressed: %v", got)
	}
}

func TestReviewedDropsExposeNoIdentityValues(t *testing.T) {
	output, _ := renderFixtureMetrics(t, "primary")
	for _, forbidden := range []string{"mongo-1.example.invalid:27017", "rs_fixture", "config_fixture/", "CN=mongodb.example.invalid"} {
		if strings.Contains(string(output), forbidden) {
			t.Errorf("reviewed identity %q exposed", forbidden)
		}
	}
}
func TestGeneratedFamilySchemasHaveUniqueSamples(t *testing.T) {
	seen := map[string]string{}
	for _, definition := range metricDefinitions {
		keys := make([]string, 0, len(definition.FixedLabels))
		for key := range definition.FixedLabels {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		var labels strings.Builder
		for _, key := range keys {
			labels.WriteString(key)
			labels.WriteByte('=')
			labels.WriteString(definition.FixedLabels[key])
			labels.WriteByte(0)
		}
		identity := definition.Family + "\x00" + labels.String()
		if previous, exists := seen[identity]; exists {
			t.Errorf("duplicate sample identity for %s and %s", previous, definition.SourcePath)
		}
		seen[identity] = definition.SourcePath
	}
}
