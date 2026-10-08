// Copyright 2017 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/percona/exporter_shared/helpers"
	"github.com/stretchr/testify/assert"

	"github.com/percona/mongodb_exporter/collector/mongod"
	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
)

func testMongoDBURL() string {
	if u := os.Getenv("TEST_MONGODB_URI"); u != "" {
		return u
	}
	return "mongodb://localhost:27017"
}

func TestCollector(t *testing.T) {
	if testing.Short() {
		t.Skip("-short is passed, skipping functional test")
	}

	collector := NewMongodbCollector(&MongodbCollectorOpts{
		URI:                      testMongoDBURL(),
		CollectDatabaseMetrics:   true,
		CollectCollectionMetrics: true,
		CollectTopMetrics:        true,
		CollectIndexUsageStats:   true,
	})

	descCh := make(chan *prometheus.Desc)
	go func() {
		collector.Describe(descCh)
		close(descCh)
	}()
	metricCh := make(chan prometheus.Metric)
	go func() {
		collector.Collect(metricCh)
		close(metricCh)
	}()

	descriptors := make(map[string]struct{})
	var descriptorsCount int
	for d := range descCh {
		descriptors[d.String()] = struct{}{}
		descriptorsCount++
	}

	var metricsCount int
	var versionInfoFound bool
	for m := range metricCh {

		if _, ok := descriptors[m.Desc().String()]; ok {
			delete(descriptors, m.Desc().String())
		}

		m := helpers.ReadMetric(m)
		switch m.Name {
		case "mongodb_version_info":
			versionInfoFound = true
		}
		metricsCount++
	}

	var missingDescMsg string
	for k := range descriptors {
		missingDescMsg += fmt.Sprintf("- %s\n", k)
	}

	assert.True(t, len(descriptors) == 0, "Number of descriptors collected and described should be the same. "+
		"Got '%d' Descriptors from collector.Describe() and '%d' from collector.Collect().\n"+
		"Missing descriptors: \n%s", descriptorsCount, metricsCount, missingDescMsg)
	assert.True(t, versionInfoFound, "version info metric not found")
}

func loadFastCollectorFixture(t *testing.T) *mongod.ServerStatus {
	t.Helper()
	source, err := os.ReadFile(filepath.Join("serverstatusv5", "testdata", "source", "server-status-primary-5.0.34.json"))
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
	status := &mongod.ServerStatus{Raw: raw}
	if err := bson.Unmarshal(raw, status); err != nil {
		t.Fatal(err)
	}
	return status
}

func collectFastServerStatusMetrics(t *testing.T, collector *MongodbCollector, status *mongod.ServerStatus) map[string][]*helpers.Metric {
	t.Helper()
	metrics := make(chan prometheus.Metric, 4096)
	collector.collectMongodServerStatus(status, metrics)
	close(metrics)
	result := make(map[string][]*helpers.Metric)
	for metric := range metrics {
		decoded := helpers.ReadMetric(metric)
		result[decoded.Name] = append(result[decoded.Name], decoded)
	}
	return result
}

func fastLegacyMetrics(metrics map[string][]*helpers.Metric) map[string][]*helpers.Metric {
	legacy := make(map[string][]*helpers.Metric)
	for family, samples := range metrics {
		if !strings.HasPrefix(family, "mongodb_server_status_") {
			legacy[family] = samples
		}
	}
	return legacy
}

func canonicalFastMetrics(metrics map[string][]*helpers.Metric) map[string][]string {
	canonical := make(map[string][]string, len(metrics))
	for family, samples := range metrics {
		canonical[family] = make([]string, 0, len(samples))
		for _, sample := range samples {
			labelNames := make([]string, 0, len(sample.Labels))
			for name := range sample.Labels {
				labelNames = append(labelNames, name)
			}
			sort.Strings(labelNames)

			var projection strings.Builder
			fmt.Fprintf(&projection, "name=%q help=%q type=%q", sample.Name, sample.Help, sample.Type.String())
			for _, name := range labelNames {
				fmt.Fprintf(&projection, " label[%q]=%q", name, sample.Labels[name])
			}
			fmt.Fprintf(&projection, " value_bits=%016x", math.Float64bits(sample.Value))
			canonical[family] = append(canonical[family], projection.String())
		}
		sort.Strings(canonical[family])
	}
	return canonical
}

func assertSameFastMetrics(t *testing.T, want, got map[string][]*helpers.Metric) {
	t.Helper()
	if !assert.Equal(t, canonicalFastMetrics(want), canonicalFastMetrics(got)) {
		t.Error("legacy fast metric names, HELP, TYPE, labels, or values changed under modern suppression")
	}
}

func TestFastCollectorUsesScopedWiredTigerModule(t *testing.T) {
	status := loadFastCollectorFixture(t)
	metrics := collectFastServerStatusMetrics(t, NewMongodbCollector(&MongodbCollectorOpts{}), status)

	for _, family := range []string{
		"mongodb_version_info",
		"mongodb_mongod_wiredtiger_cache_bytes",
		"mongodb_server_status_wired_tiger_cache_bytes_allocated_for_updates_total",
	} {
		if len(metrics[family]) == 0 {
			t.Errorf("fast collector omitted %s", family)
		}
	}
	for family := range metrics {
		if strings.HasPrefix(family, "mongodb_server_status_") &&
			!strings.HasPrefix(family, "mongodb_server_status_wired_tiger_") &&
			family != "mongodb_server_status_decode_errors_total" {
			t.Errorf("fast collector emitted non-WiredTiger modern family %s", family)
		}
	}
	if len(metrics["mongodb_server_status_commands_total"]) != 0 {
		t.Error("fast collector emitted detailed command metrics")
	}
}

func TestSuppressServerStatusV5KeepsLegacyFastMetrics(t *testing.T) {
	status := loadFastCollectorFixture(t)
	enabled := collectFastServerStatusMetrics(t, NewMongodbCollector(&MongodbCollectorOpts{}), status)
	suppressed := collectFastServerStatusMetrics(t, NewMongodbCollector(&MongodbCollectorOpts{SuppressServerStatusV5: true}), status)

	for _, family := range []string{
		"mongodb_version_info",
		"mongodb_mongod_wiredtiger_cache_bytes",
	} {
		if len(suppressed[family]) == 0 {
			t.Errorf("suppression removed legacy fast family %s", family)
		}
	}
	assertSameFastMetrics(t, fastLegacyMetrics(enabled), suppressed)
	for family := range suppressed {
		if strings.HasPrefix(family, "mongodb_server_status_") {
			t.Errorf("suppression retained modern family %s", family)
		}
	}
}
