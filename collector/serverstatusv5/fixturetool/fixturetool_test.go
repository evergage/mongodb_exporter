package fixturetool

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	redactionVersion               = "2"
	reconstructionVersion          = "2"
	normalizationVersion           = "2"
	originalBSONTypesAvailable     = false
	fixtureUnionScalarLeafCount    = 1904
	fixtureAddedCommandPathCount   = 323
	primaryPreExtensionLeafCount   = 1573
	secondaryPreExtensionLeafCount = 1502
)

var fixtureMetadata = []struct {
	role, mongoVersion, captureTime, sourceSHA256, reconstructedSHA256 string
}{
	{"primary", "5.0.34", "2026-09-21T21:05:28.899Z", "347d24bcc8716a81673403ef17ad2d13bdf85cd5b0453653ab5c08fbad452e2a", "6024a88e3f993fc2bf41381239a2affc2c0420d012d18f1a298580bfffd291fe"},
	{"secondary", "5.0.34", "2026-09-21T20:42:39.325Z", "125d6088d2553962eaa1724c0d626f5884bed2d3dfb50c4201ba6832a0659f10", "0b216ddb4b8da5b54f4e563dbf5745eb63220322bef5203f72f9f23845663502"},
}

func readFixture(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var data bytes.Buffer
	_, err = data.ReadFrom(file)
	return data.Bytes(), err
}

func TestFixtureProvenanceAndDeterministicReconstruction(t *testing.T) {
	if redactionVersion == "" || reconstructionVersion == "" || normalizationVersion == "" {
		t.Fatal("fixture pipeline versions must be recorded")
	}
	if originalBSONTypesAvailable {
		t.Fatal("fixtures must record that source BSON types are unavailable")
	}
	for _, metadata := range fixtureMetadata {
		t.Run(metadata.role, func(t *testing.T) {
			if metadata.mongoVersion == "" || metadata.captureTime == "" {
				t.Fatal("fixture role, version, and capture time must be recorded")
			}
			source, err := readFixture(filepath.Join("..", "testdata", "source", "server-status-"+metadata.role+"-5.0.34.json"))
			if err != nil {
				t.Fatal(err)
			}
			if err := ValidateSanitizedJSON(source); err != nil {
				t.Fatal(err)
			}
			if got := fmt.Sprintf("%x", sha256.Sum256(source)); got != metadata.sourceSHA256 {
				t.Fatalf("source digest=%s want %s", got, metadata.sourceSHA256)
			}
			first, err := ReconstructJSON(source)
			if err != nil {
				t.Fatal(err)
			}
			if got := fmt.Sprintf("%x", sha256.Sum256(first)); got != metadata.reconstructedSHA256 {
				t.Fatalf("reconstructed digest=%s want %s", got, metadata.reconstructedSHA256)
			}
			second, err := ReconstructJSON(source)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(first, second) {
				t.Fatal("reconstruction is not deterministic")
			}
		})
	}
}

func TestFixtureUnionHasExpectedScalarLeafCount(t *testing.T) {
	union := make(map[string]struct{})
	priorCounts := map[string]int{"primary": primaryPreExtensionLeafCount, "secondary": secondaryPreExtensionLeafCount}
	for _, metadata := range fixtureMetadata {
		source, err := readFixture(filepath.Join("..", "testdata", "source", "server-status-"+metadata.role+"-5.0.34.json"))
		if err != nil {
			t.Fatal(err)
		}
		fixture, err := decodeJSON(source)
		if err != nil {
			t.Fatal(err)
		}
		paths := SortedLeafPaths(fixture)
		if got := len(paths) - priorCounts[metadata.role]; got != fixtureAddedCommandPathCount {
			t.Fatalf("%s added scalar leaves=%d want %d", metadata.role, got, fixtureAddedCommandPathCount)
		}
		for _, path := range paths {
			union[path] = struct{}{}
		}
	}
	if got := len(union); got != fixtureUnionScalarLeafCount {
		t.Fatalf("fixture union scalar leaves=%d want %d", got, fixtureUnionScalarLeafCount)
	}
}

func TestSanitizationGuardsRejectCapturedValues(t *testing.T) {
	probes := []string{
		`{"host":"captured-host"}`,
		`{"repl":{"setName":"captured-set"}}`,
		`{"sharding":{"configsvrConnectionString":"captured/captured:27017"}}`,
		`{"pid":1234}`,
		`{"metrics":{"commands":{"find":{"total":99}}}}`,
	}
	for _, probe := range probes {
		if err := ValidateSanitizedJSON([]byte(probe)); err == nil {
			t.Errorf("captured value accepted: %s", probe)
		}
	}
}

func TestSanitizePreservesShapeAndExtendedBSONPaths(t *testing.T) {
	source := []byte(`{"host":"captured","pid":4321,"enabled":true,"$clusterTime":{"clusterTime":[12,3],"signature":{"hash":"captured","keyId":7}},"repl":{"topologyVersion":{"processId":"123456789012345678901234"}}}`)
	sanitized, err := SanitizeJSON(source)
	if err != nil {
		t.Fatal(err)
	}
	if err := ValidateSanitizedJSON(sanitized); err != nil {
		t.Fatal(err)
	}
	value, err := decodeJSON(sanitized)
	if err != nil {
		t.Fatal(err)
	}
	clusterTime := value.(map[string]interface{})["$clusterTime"].(map[string]interface{})["clusterTime"].([]interface{})
	if got := clusterTime[0].(json.Number).String(); got != "1577934245" {
		t.Errorf("sanitized cluster time seconds=%s want 1577934245", got)
	}
	if got := clusterTime[1].(json.Number).String(); got != "7" {
		t.Errorf("sanitized cluster time increment=%s want 7", got)
	}
	if got := len(SortedLeafPaths(value)); got != 8 {
		t.Fatalf("sanitized scalar leaves=%d want 8", got)
	}
	reconstructed, err := ReconstructJSON(sanitized)
	if err != nil {
		t.Fatal(err)
	}
	for _, wrapper := range [][]byte{[]byte(`"$timestamp"`), []byte(`"$binary"`), []byte(`"$oid"`)} {
		if !bytes.Contains(reconstructed, wrapper) {
			t.Errorf("reconstruction lacks %s", wrapper)
		}
	}
}

func TestSnapshotWindowStringsSanitizeWithoutCapturedValues(t *testing.T) {
	source := []byte(`{"wiredTiger":{"snapshot-window-settings":{"latest majority snapshot timestamp available":"Sep 21 21:05:28:123","oldest majority snapshot timestamp available":"Sep 21 21:00:28:123"}}}`)
	sanitized, err := SanitizeJSON(source)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(sanitized, []byte("Sep 21")) {
		t.Fatalf("sanitized fixture retained captured timestamp: %s", sanitized)
	}
	for _, expected := range [][]byte{[]byte(`"Jan 02 03:04:05:7"`), []byte(`"Jan 02 02:59:05:7"`)} {
		if !bytes.Contains(sanitized, expected) {
			t.Errorf("sanitized fixture lacks synthetic snapshot %s", expected)
		}
	}

	reconstructed, err := ReconstructJSON(sanitized)
	if err != nil {
		t.Fatal(err)
	}
	value, err := decodeJSON(reconstructed)
	if err != nil {
		t.Fatal(err)
	}
	snapshotSettings := value.(map[string]interface{})["wiredTiger"].(map[string]interface{})["snapshot-window-settings"].(map[string]interface{})
	for _, path := range []string{"latest majority snapshot timestamp available", "oldest majority snapshot timestamp available"} {
		if got := snapshotSettings[path].(string); !strings.HasPrefix(got, "Jan 02 ") || !strings.HasSuffix(got, ":7") {
			t.Errorf("%s=%q want synthetic snapshot string", path, got)
		}
	}
}
