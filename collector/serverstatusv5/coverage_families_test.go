package serverstatusv5

import (
	"io/ioutil"
	"math"
	"path/filepath"
	"testing"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"go.mongodb.org/mongo-driver/bson"
)

type rawModuleCollector struct {
	module *Module
	raw    bson.Raw
}

func (c *rawModuleCollector) Describe(ch chan<- *prometheus.Desc) { c.module.Describe(ch) }
func (c *rawModuleCollector) Collect(ch chan<- prometheus.Metric) { c.module.Collect(c.raw, ch) }

func loadFixtureRaw(t *testing.T, role string) bson.Raw {
	t.Helper()
	source, err := ioutil.ReadFile(filepath.Join("testdata", "source", "server-status-"+role+"-5.0.34.json"))
	if err != nil {
		t.Fatal(err)
	}
	data, err := fixturetool.ReconstructJSON(source)
	if err != nil {
		t.Fatal(err)
	}
	var raw bson.Raw
	if err := bson.UnmarshalExtJSON(data, true, &raw); err != nil {
		t.Fatal(err)
	}
	return raw
}
func gatherFixture(t *testing.T, role string) map[string]*dto.MetricFamily {
	t.Helper()
	return gatherFixtureRaw(t, loadFixtureRaw(t, role))
}

func gatherFixtureRaw(t *testing.T, raw bson.Raw) map[string]*dto.MetricFamily {
	t.Helper()
	collector := &rawModuleCollector{module: New(), raw: raw}
	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(collector); err != nil {
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

func fixtureScalar(t *testing.T, raw bson.Raw, path ...string) float64 {
	t.Helper()
	value := raw.Lookup(path...)
	if actual, ok := value.DoubleOK(); ok {
		return actual
	}
	if actual, ok := value.Int32OK(); ok {
		return float64(actual)
	}
	if actual, ok := value.Int64OK(); ok {
		return float64(actual)
	}
	if actual, ok := value.BooleanOK(); ok {
		if actual {
			return 1
		}
		return 0
	}
	if actual, ok := value.DateTimeOK(); ok {
		return float64(actual)
	}
	t.Fatalf("fixture path %v has unsupported BSON type %s", path, value.Type)
	return 0
}

func assertFixtureMetricValue(t *testing.T, families map[string]*dto.MetricFamily, raw bson.Raw, family string, labels map[string]string, conversion float64, path ...string) {
	t.Helper()
	metric := metricWithLabels(families[family], labels)
	if metric == nil {
		t.Errorf("%s%v missing", family, labels)
		return
	}
	want := fixtureScalar(t, raw, path...) * conversion
	got := metricValueDTO(metric)
	tolerance := math.Max(1e-12, math.Abs(want)*1e-12)
	if math.Abs(got-want) > tolerance {
		t.Errorf("%s%v=%v want %v from fixture path %v with conversion %v", family, labels, got, want, path, conversion)
	}
}
func metricWithLabels(family *dto.MetricFamily, labels map[string]string) *dto.Metric {
	if family == nil {
		return nil
	}
	for _, metric := range family.Metric {
		actual := map[string]string{}
		for _, label := range metric.Label {
			actual[label.GetName()] = label.GetValue()
		}
		match := true
		for key, value := range labels {
			if actual[key] != value {
				match = false
				break
			}
		}
		if match {
			return metric
		}
	}
	return nil
}
func metricValueDTO(metric *dto.Metric) float64 {
	if metric == nil {
		return 0
	}
	if metric.Gauge != nil {
		return metric.Gauge.GetValue()
	}
	if metric.Counter != nil {
		return metric.Counter.GetValue()
	}
	return 0
}

func TestRoleElectionAndPrimaryOnlyServiceCoverage(t *testing.T) {
	primary := gatherFixture(t, "primary")
	secondary := gatherFixture(t, "secondary")
	roleFamily := "mongodb_server_status_repl_role"
	if got := metricValueDTO(metricWithLabels(primary[roleFamily], map[string]string{"state": "primary"})); got != 1 {
		t.Errorf("primary role=%v", got)
	}
	if got := metricValueDTO(metricWithLabels(secondary[roleFamily], map[string]string{"state": "primary"})); got != 0 {
		t.Errorf("secondary primary-role=%v", got)
	}
	if got := metricValueDTO(metricWithLabels(secondary[roleFamily], map[string]string{"state": "secondary"})); got != 1 {
		t.Errorf("secondary role=%v", got)
	}
	election := primary["mongodb_server_status_election_calls_total"]
	if metricWithLabels(election, map[string]string{"reason": "step_up_cmd", "outcome": "successful"}) == nil {
		t.Error("primary step-up success metric missing")
	}
	service := "mongodb_server_status_repl_primary_only_service_state"
	if metricWithLabels(primary[service], map[string]string{"service": "resharding_donor_service", "state": "running"}) == nil {
		t.Error("primary running service state missing")
	}
	if metricWithLabels(secondary[service], map[string]string{"service": "resharding_donor_service", "state": "paused"}) == nil {
		t.Error("secondary paused service state missing")
	}
}
func TestFlowControlCoverageAndConversions(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	checks := []struct {
		family     string
		conversion float64
		path       []string
	}{
		{"mongodb_server_status_flow_control_enabled", 1, []string{"flowControl", "enabled"}},
		{"mongodb_server_status_flow_control_is_lagged", 1, []string{"flowControl", "isLagged"}},
		{"mongodb_server_status_flow_control_is_lagged_count_total", 1, []string{"flowControl", "isLaggedCount"}},
		{"mongodb_server_status_flow_control_is_lagged_time_seconds", 1e-6, []string{"flowControl", "isLaggedTimeMicros"}},
		{"mongodb_server_status_flow_control_time_acquiring_seconds", 1e-6, []string{"flowControl", "timeAcquiringMicros"}},
		{"mongodb_server_status_flow_control_sustainer_rate", 1, []string{"flowControl", "sustainerRate"}},
		{"mongodb_server_status_flow_control_target_rate_limit", 1, []string{"flowControl", "targetRateLimit"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, check.family, nil, check.conversion, check.path...)
	}
}
func TestShardingStatisticsCoverage(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	checks := []struct {
		labels map[string]string
		path   []string
	}{
		{map[string]string{"direction": "donor", "outcome": "started"}, []string{"shardingStatistics", "countDonorMoveChunkStarted"}},
		{map[string]string{"direction": "donor", "outcome": "committed"}, []string{"shardingStatistics", "countDonorMoveChunkCommitted"}},
		{map[string]string{"direction": "donor", "outcome": "aborted"}, []string{"shardingStatistics", "countDonorMoveChunkAborted"}},
		{map[string]string{"direction": "recipient", "outcome": "started"}, []string{"shardingStatistics", "countRecipientMoveChunkStarted"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, "mongodb_server_status_sharding_move_chunk_total", check.labels, 1, check.path...)
	}
	static := []struct {
		family string
		path   []string
	}{
		{"mongodb_server_status_sharding_statistics_count_bytes_cloned_on_recipient_total", []string{"shardingStatistics", "countBytesClonedOnRecipient"}},
		{"mongodb_server_status_sharding_statistics_count_bytes_cloned_on_donor_total", []string{"shardingStatistics", "countBytesClonedOnDonor"}},
		{"mongodb_server_status_sharding_statistics_count_stale_config_errors_total", []string{"shardingStatistics", "countStaleConfigErrors"}},
	}
	for _, check := range static {
		assertFixtureMetricValue(t, primary, raw, check.family, nil, 1, check.path...)
	}
}
func TestSessionTransactionReadConcernAndOplogCoverage(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	checks := []struct {
		family     string
		conversion float64
		path       []string
	}{
		{"mongodb_server_status_logical_session_record_cache_active_sessions_count", 1, []string{"logicalSessionRecordCache", "activeSessionsCount"}},
		{"mongodb_server_status_logical_session_record_cache_last_sessions_collection_job_duration_seconds", 1e-3, []string{"logicalSessionRecordCache", "lastSessionsCollectionJobDurationMillis"}},
		{"mongodb_server_status_transactions_current_open", 1, []string{"transactions", "currentOpen"}},
		{"mongodb_server_status_transactions_total_committed_total", 1, []string{"transactions", "totalCommitted"}},
		{"mongodb_server_status_read_concern_counters_non_transaction_ops_majority_total", 1, []string{"readConcernCounters", "nonTransactionOps", "majority"}},
		{"mongodb_server_status_oplog_truncation_total_time_truncating_seconds", 1e-6, []string{"oplogTruncation", "totalTimeTruncatingMicros"}},
		{"mongodb_server_status_oplog_truncation_truncate_count_total", 1, []string{"oplogTruncation", "truncateCount"}},
		{"mongodb_server_status_repl_rbid", 1, []string{"repl", "rbid"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, check.family, nil, check.conversion, check.path...)
	}
	var sparse bson.Raw
	sparseBytes, err := bson.Marshal(bson.M{"transactions": bson.M{"currentOpen": 0}})
	if err != nil {
		t.Fatal(err)
	}
	sparse = bson.Raw(sparseBytes)
	metrics := make(chan prometheus.Metric, 32)
	report := New().Collect(sparse, metrics)
	close(metrics)
	if len(report.Errors) > 0 {
		t.Fatalf("sparse errors: %+v", report.Errors)
	}
	for metric := range metrics {
		if metric.Desc().String() == "" {
			t.Fatal("invalid metric")
		}
	}
}
func TestBoundedWorkloadCoverage(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	checks := []struct {
		family string
		labels map[string]string
		path   []string
	}{
		{"mongodb_server_status_commands_total", map[string]string{"command": "find", "outcome": "total"}, []string{"metrics", "commands", "find", "total"}},
		{"mongodb_server_status_commands_total", map[string]string{"command": "find", "outcome": "failed"}, []string{"metrics", "commands", "find", "failed"}},
		{"mongodb_server_status_aggregation_stage_total", map[string]string{"stage": "$group"}, []string{"metrics", "aggStageCounters", "$group"}},
		{"mongodb_server_status_operator_total", map[string]string{"category": "expressions", "operator": "$eq"}, []string{"metrics", "operatorCounters", "expressions", "$eq"}},
		{"mongodb_server_status_metrics_query_update_many_count_total", nil, []string{"metrics", "query", "updateManyCount"}},
		{"mongodb_server_status_metrics_cursor_total_opened_total", nil, []string{"metrics", "cursor", "totalOpened"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, check.family, check.labels, 1, check.path...)
	}
}
func TestNetworkSecurityAndStorageCoverage(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	checks := []struct {
		family     string
		labels     map[string]string
		conversion float64
		path       []string
	}{
		{"mongodb_server_status_network_physical_bytes_in_total", nil, 1, []string{"network", "physicalBytesIn"}},
		{"mongodb_server_status_network_num_slow_ssloperations_total", nil, 1, []string{"network", "numSlowSSLOperations"}},
		{"mongodb_server_status_network_tcp_fast_open_server_supported", nil, 1, []string{"network", "tcpFastOpen", "serverSupported"}},
		{"mongodb_server_status_security_sslserver_certificate_expiration_date_seconds", nil, 1e-3, []string{"security", "SSLServerCertificateExpirationDate"}},
		{"mongodb_server_status_storage_engine_supports_committed_reads", nil, 1, []string{"storageEngine", "supportsCommittedReads"}},
		{"mongodb_server_status_storage_engine_drop_pending_idents", nil, 1, []string{"storageEngine", "dropPendingIdents"}},
		{"mongodb_server_status_network_compression_bytes_total", map[string]string{"compressor": "zstd", "operation": "compressor", "direction": "in"}, 1, []string{"network", "compression", "zstd", "compressor", "bytesIn"}},
		{"mongodb_server_status_network_service_executor", map[string]string{"executor": "passthrough", "state": "threads_running"}, 1, []string{"network", "serviceExecutors", "passthrough", "threadsRunning"}},
		{"mongodb_server_status_security_authentication_total", map[string]string{"mechanism": "MONGODB-X509", "conversation": "authenticate", "outcome": "successful"}, 1, []string{"security", "authentication", "mechanisms", "MONGODB-X509", "authenticate", "successful"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, check.family, check.labels, check.conversion, check.path...)
	}
}
func TestSelectedWiredTigerCoverage(t *testing.T) {
	raw := loadFixtureRaw(t, "primary")
	primary := gatherFixtureRaw(t, raw)
	for _, family := range []string{"mongodb_server_status_wired_tiger_concurrent_transactions", "mongodb_server_status_metrics_cursor_open"} {
		if primary[family] != nil {
			t.Errorf("duplicate modern family %s remains", family)
		}
	}
	checks := []struct {
		family string
		path   []string
	}{
		{"mongodb_server_status_wired_tiger_snapshot_window_settings_total_number_of_snapshot_too_old_errors_total", []string{"wiredTiger", "snapshot-window-settings", "total number of SnapshotTooOld errors"}},
		{"mongodb_server_status_wired_tiger_snapshot_window_settings_current_available_snapshot_window_size_in_seconds", []string{"wiredTiger", "snapshot-window-settings", "current available snapshot window size in seconds"}},
	}
	for _, check := range checks {
		assertFixtureMetricValue(t, primary, raw, check.family, nil, 1, check.path...)
	}
	contract, paths := loadCoverageInputs(t)
	if err := contract.Audit(paths); err != nil {
		t.Fatal(err)
	}
}
