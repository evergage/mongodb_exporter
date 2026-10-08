package serverstatusv5

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"go.mongodb.org/mongo-driver/bson"
)

func gatherScopedModule(t *testing.T, module *Module, raw bson.Raw) map[string]*dto.MetricFamily {
	t.Helper()
	collector := &rawModuleCollector{module: module, raw: raw}
	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(collector); err != nil {
		t.Fatal(err)
	}
	return gatherRegistered(t, registry)
}

func gatherScopedFixture(t *testing.T, module *Module, role string) map[string]*dto.MetricFamily {
	t.Helper()
	return gatherScopedModule(t, module, loadFixtureRaw(t, role))
}

func scopedDescriptorFamilies(module *Module) map[string]struct{} {
	families := make(map[string]struct{}, len(module.descriptors))
	for family := range module.descriptors {
		families[family] = struct{}{}
	}
	return families
}

func scopedMetricIdentity(family string, metric *dto.Metric) string {
	labels := make([]string, 0, len(metric.GetLabel()))
	for _, label := range metric.GetLabel() {
		labels = append(labels, strconv.Quote(label.GetName())+"="+strconv.Quote(label.GetValue()))
	}
	sort.Strings(labels)
	return strconv.Quote(family) + "{" + strings.Join(labels, ",") + "}"
}

func scopedMetricIdentities(t *testing.T, families map[string]*dto.MetricFamily) map[string]struct{} {
	t.Helper()
	identities := map[string]struct{}{}
	for family, metrics := range families {
		for _, metric := range metrics.GetMetric() {
			identity := scopedMetricIdentity(family, metric)
			if _, exists := identities[identity]; exists {
				t.Fatalf("duplicate sample identity %s", identity)
			}
			identities[identity] = struct{}{}
		}
	}
	return identities
}

func mergeScopedSets(sets ...map[string]struct{}) map[string]struct{} {
	merged := map[string]struct{}{}
	for _, set := range sets {
		for item := range set {
			merged[item] = struct{}{}
		}
	}
	return merged
}

func assertScopedSetsDisjoint(t *testing.T, left, right map[string]struct{}, kind string) {
	t.Helper()
	for item := range left {
		if _, exists := right[item]; exists {
			t.Fatalf("%s crosses scopes: %s", kind, item)
		}
	}
}

func TestScopedDescriptorsAndFixtureIdentitiesAreDisjoint(t *testing.T) {
	wiredTigerDescriptors := scopedDescriptorFamilies(NewWiredTiger())
	detailedDescriptors := scopedDescriptorFamilies(NewDetailed())
	assertScopedSetsDisjoint(t, wiredTigerDescriptors, detailedDescriptors, "descriptor family")
	if got, want := mergeScopedSets(wiredTigerDescriptors, detailedDescriptors), scopedDescriptorFamilies(New()); !reflect.DeepEqual(got, want) {
		t.Fatalf("scoped descriptor union has %d families, complete module has %d", len(got), len(want))
	}

	for _, role := range []string{"primary", "secondary"} {
		wiredTiger := scopedMetricIdentities(t, gatherScopedFixture(t, NewWiredTiger(), role))
		detailed := scopedMetricIdentities(t, gatherScopedFixture(t, NewDetailed(), role))
		complete := scopedMetricIdentities(t, gatherScopedFixture(t, New(), role))
		assertScopedSetsDisjoint(t, wiredTiger, detailed, role+" sample identity")
		if got, want := mergeScopedSets(wiredTiger, detailed), complete; !reflect.DeepEqual(got, want) {
			t.Fatalf("%s scoped union has %d identities, complete module has %d", role, len(got), len(want))
		}
	}
}

func TestWiredTigerScopeEmitsAllFixtureMappings(t *testing.T) {
	const expectedMappings = 495

	for _, role := range []string{"primary", "secondary"} {
		module := NewWiredTiger()
		if got := len(module.definitions); got != expectedMappings {
			t.Fatalf("%s has %d WiredTiger definitions, want %d", role, got, expectedMappings)
		}
		for _, definition := range module.definitions {
			if len(definition.PathSegments) == 0 || definition.PathSegments[0] != "wiredTiger" {
				t.Fatalf("%s includes non-WiredTiger definition %q", role, definition.SourcePath)
			}
		}

		metrics := make(chan prometheus.Metric, len(metricDefinitions)+len(generatedDynamicFamilies)+1)
		report := module.Collect(loadFixtureRaw(t, role), metrics)
		close(metrics)
		if len(report.Errors) != 0 {
			t.Fatalf("%s WiredTiger decode errors: %+v", role, report.Errors)
		}
		if report.Emitted != expectedMappings {
			t.Fatalf("%s emitted %d WiredTiger mappings, want %d", role, report.Emitted, expectedMappings)
		}
		emitted := 0
		for range metrics {
			emitted++
		}
		if emitted != expectedMappings {
			t.Fatalf("%s exported %d WiredTiger samples, want %d", role, emitted, expectedMappings)
		}
	}
}

func numericDefinitionForScope(t *testing.T, scope moduleScope) metricDefinition {
	t.Helper()
	for _, definition := range metricDefinitions {
		if definitionInScope(definition, scope) && definition.Transform == "" && definition.ValueLabel == "" {
			return definition
		}
	}
	t.Fatalf("no numeric definition for scope %d", scope)
	return metricDefinition{}
}

func rawWithInvalidDefinitions(t *testing.T, definitions ...metricDefinition) bson.Raw {
	t.Helper()
	document := bson.M{}
	for _, definition := range definitions {
		current := document
		for _, segment := range definition.PathSegments[:len(definition.PathSegments)-1] {
			next, ok := current[segment].(bson.M)
			if !ok {
				next = bson.M{}
				current[segment] = next
			}
			current = next
		}
		current[definition.PathSegments[len(definition.PathSegments)-1]] = "not numeric"
	}
	data, err := bson.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	return bson.Raw(data)
}

func collectScopedReport(module *Module, raw bson.Raw) Report {
	metrics := make(chan prometheus.Metric, len(metricDefinitions)+len(generatedDynamicFamilies)+1)
	report := module.Collect(raw, metrics)
	close(metrics)
	return report
}

func reportIncludesFamily(report Report, family string) bool {
	for _, failure := range report.Errors {
		if failure.Family == family {
			return true
		}
	}
	return false
}

func requiredScopedMetric(t *testing.T, families map[string]*dto.MetricFamily, family string, labels map[string]string) *dto.Metric {
	t.Helper()
	metric := metricWithLabels(families[family], labels)
	if metric == nil {
		t.Fatalf("missing metric %s with labels %v", family, labels)
	}
	return metric
}

func TestScopedDiagnosticsFollowDefinitions(t *testing.T) {
	wiredTiger := numericDefinitionForScope(t, moduleScopeWiredTiger)
	detailed := numericDefinitionForScope(t, moduleScopeDetailed)
	raw := rawWithInvalidDefinitions(t, wiredTiger, detailed)

	fastReport := collectScopedReport(NewWiredTiger(), raw)
	if len(fastReport.Errors) != 1 || fastReport.Errors[0].Family != wiredTiger.Family {
		t.Fatalf("fast decode errors=%+v, want only %s", fastReport.Errors, wiredTiger.Family)
	}
	detailReport := collectScopedReport(NewDetailed(), raw)
	if len(detailReport.Errors) != 1 || detailReport.Errors[0].Family != detailed.Family {
		t.Fatalf("detail decode errors=%+v, want only %s", detailReport.Errors, detailed.Family)
	}
	completeReport := collectScopedReport(New(), raw)
	if !reportIncludesFamily(completeReport, wiredTiger.Family) || !reportIncludesFamily(completeReport, detailed.Family) {
		t.Fatalf("complete module lost decode diagnostics: %+v", completeReport.Errors)
	}

	data, err := bson.Marshal(bson.M{"metrics": bson.M{"commands": bson.M{"unreviewed": bson.M{"total": 1}}}})
	if err != nil {
		t.Fatal(err)
	}
	unknownRaw := bson.Raw(data)
	fastFamilies := gatherScopedModule(t, NewWiredTiger(), unknownRaw)
	if _, exists := fastFamilies["mongodb_server_status_unknown_keys_total"]; exists {
		t.Fatal("fast scope emitted unknown-key diagnostics")
	}
	detailFamilies := gatherScopedModule(t, NewDetailed(), unknownRaw)
	unknown := requiredScopedMetric(t, detailFamilies, "mongodb_server_status_unknown_keys_total", map[string]string{"family": "commands"})
	if got := metricValueDTO(unknown); got != 1 {
		t.Fatalf("detail unknown-key diagnostic=%v, want 1", got)
	}
	completeFamilies := gatherScopedModule(t, New(), unknownRaw)
	unknown = requiredScopedMetric(t, completeFamilies, "mongodb_server_status_unknown_keys_total", map[string]string{"family": "commands"})
	if got := metricValueDTO(unknown); got != 1 {
		t.Fatalf("complete unknown-key diagnostic=%v, want 1", got)
	}
}

func TestScopedModulesDoNotLeakRoleTransitions(t *testing.T) {
	const wiredTigerFamily = "mongodb_server_status_wired_tiger_cache_bytes_allocated_for_updates_total"
	const commandsFamily = "mongodb_server_status_commands_total"
	roles := []string{"primary", "secondary"}

	wiredTigerPresence := map[string]bool{"primary": true, "secondary": false}
	rawForRole := func(role string) bson.Raw {
		document := bson.M{}
		if wiredTigerPresence[role] {
			document["wiredTiger"] = bson.M{"cache": bson.M{"bytes allocated for updates": int64(1)}}
		}
		data, err := bson.Marshal(document)
		if err != nil {
			t.Fatal(err)
		}
		return bson.Raw(data)
	}
	for _, sequence := range [][]string{{"primary", "secondary"}, {"secondary", "primary"}} {
		t.Run("wired-tiger-"+sequence[0]+"-to-"+sequence[1], func(t *testing.T) {
			collector := &rawModuleCollector{module: NewWiredTiger(), raw: rawForRole(sequence[0])}
			registry := prometheus.NewPedanticRegistry()
			if err := registry.Register(collector); err != nil {
				t.Fatal(err)
			}
			first := gatherRegistered(t, registry)
			firstPresent := metricWithLabels(first[wiredTigerFamily], nil) != nil
			if firstPresent != wiredTigerPresence[sequence[0]] {
				t.Fatalf("first WiredTiger presence=%t, want %t", firstPresent, wiredTigerPresence[sequence[0]])
			}
			collector.raw = rawForRole(sequence[1])
			second := gatherRegistered(t, registry)
			secondPresent := metricWithLabels(second[wiredTigerFamily], nil) != nil
			if secondPresent != wiredTigerPresence[sequence[1]] {
				t.Fatalf("transitioned WiredTiger presence=%t, want %t", secondPresent, wiredTigerPresence[sequence[1]])
			}
		})
	}

	detailedPresence := map[string]bool{}
	for _, role := range roles {
		families := gatherScopedFixture(t, NewDetailed(), role)
		detailedPresence[role] = metricWithLabels(families[commandsFamily], map[string]string{"command": "insert", "outcome": "total"}) != nil
	}
	if detailedPresence["primary"] == detailedPresence["secondary"] {
		t.Fatal("detailed transition fixture presence must differ")
	}
	for _, sequence := range [][]string{{"primary", "secondary"}, {"secondary", "primary"}} {
		t.Run("detailed-"+sequence[0]+"-to-"+sequence[1], func(t *testing.T) {
			collector := &rawModuleCollector{module: NewDetailed(), raw: loadFixtureRaw(t, sequence[0])}
			registry := prometheus.NewPedanticRegistry()
			if err := registry.Register(collector); err != nil {
				t.Fatal(err)
			}
			first := gatherRegistered(t, registry)
			firstPresent := metricWithLabels(first[commandsFamily], map[string]string{"command": "insert", "outcome": "total"}) != nil
			if firstPresent != detailedPresence[sequence[0]] {
				t.Fatalf("first detailed insert presence=%t, want %t", firstPresent, detailedPresence[sequence[0]])
			}
			collector.raw = loadFixtureRaw(t, sequence[1])
			second := gatherRegistered(t, registry)
			secondPresent := metricWithLabels(second[commandsFamily], map[string]string{"command": "insert", "outcome": "total"}) != nil
			if secondPresent != detailedPresence[sequence[1]] {
				t.Fatalf("transitioned detailed insert presence=%t, want %t", secondPresent, detailedPresence[sequence[1]])
			}
		})
	}
}
