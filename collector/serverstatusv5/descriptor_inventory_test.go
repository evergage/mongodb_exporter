package serverstatusv5

import (
	"regexp"
	"testing"

	"github.com/percona/mongodb_exporter/collector/mongod"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
)

var descriptorNamePattern = regexp.MustCompile(`fqName: "([^"]+)"`)

func TestModernDescriptorsDoNotCollideWithLegacyDescriptors(t *testing.T) {
	legacy := map[string]bool{}
	for _, role := range []string{"primary", "secondary"} {
		var status mongod.ServerStatus
		if err := bson.Unmarshal(loadFixtureRaw(t, role), &status); err != nil {
			t.Fatal(err)
		}
		metrics := make(chan prometheus.Metric, 10000)
		status.Export(metrics)
		close(metrics)
		for metric := range metrics {
			match := descriptorNamePattern.FindStringSubmatch(metric.Desc().String())
			if len(match) != 2 {
				t.Fatalf("cannot parse descriptor %s", metric.Desc())
			}
			legacy[match[1]] = true
		}
	}
	for _, definition := range metricDefinitions {
		if legacy[definition.Family] {
			t.Errorf("modern descriptor %s for %s collides with legacy descriptor", definition.Family, definition.SourcePath)
		}
	}
}
