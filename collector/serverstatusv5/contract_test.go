package serverstatusv5

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
)

func loadCoverageInputs(t *testing.T) (*Contract, []CoverageLeaf) {
	t.Helper()
	contractData, err := ioutil.ReadFile("contract.json")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := ParseContract(contractData)
	if err != nil {
		t.Fatal(err)
	}
	leaves := map[string]CoverageLeaf{}
	for _, fixture := range []struct{ role, name string }{{"primary", "server-status-primary-5.0.34.json"}, {"secondary", "server-status-secondary-5.0.34.json"}} {
		data, err := ioutil.ReadFile(filepath.Join("testdata", "source", fixture.name))
		if err != nil {
			t.Fatal(err)
		}
		var value interface{}
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		decoder.UseNumber()
		if err := decoder.Decode(&value); err != nil {
			t.Fatal(err)
		}
		for _, leaf := range fixturetool.SortedLeaves(value) {
			current, exists := leaves[leaf.Path]
			if !exists {
				current = CoverageLeaf{Path: leaf.Path, PathSegments: leaf.Segments}
			}
			if !sameStrings(current.PathSegments, leaf.Segments) {
				t.Fatalf("fixture path %s has inconsistent segments", leaf.Path)
			}
			current.Roles = append(current.Roles, fixture.role)
			leaves[leaf.Path] = current
		}
	}
	result := make([]CoverageLeaf, 0, len(leaves))
	for _, leaf := range leaves {
		result = append(result, leaf)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Path < result[j].Path })
	return contract, result
}

func TestCoverageContractClassifiesFixtureUnionExactlyOnce(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	if err := contract.Audit(leaves); err != nil {
		t.Fatal(err)
	}
}
func TestCoverageAuditRejectsMissingClassification(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	contract.Entries = contract.Entries[1:]
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "missing=") {
		t.Fatalf("got %v", err)
	}
}
func TestCoverageAuditRejectsDuplicateClassification(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	contract.Entries = append(contract.Entries, contract.Entries[0])
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "duplicate=") {
		t.Fatalf("got %v", err)
	}
}
func TestCoverageAuditRejectsStaleClassification(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	contract.Entries = append(contract.Entries, ContractEntry{Path: "not.present", PathSegments: []string{"not", "present"}, Classification: "drop", Reason: "reviewed", DecisionSource: "test", Reviewed: true})
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "stale=") {
		t.Fatalf("got %v", err)
	}
}
func TestCoverageAuditRejectsUnreviewedDrop(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	contract.Entries[0].Classification = "drop"
	contract.Entries[0].Metric = nil
	contract.Entries[0].Reviewed = false
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "invalid=") {
		t.Fatalf("got %v", err)
	}
}
func TestCoverageAuditRejectsFixtureFactDrift(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	contract.Entries[0].PathSegments = []string{"wrong"}
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "invalid=") {
		t.Fatalf("path segments: %v", err)
	}
	contract, leaves = loadCoverageInputs(t)
	contract.Entries[0].Roles = []string{"secondary"}
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "invalid=") {
		t.Fatalf("roles: %v", err)
	}
}
func TestCoverageAuditRejectsInvalidConsumer(t *testing.T) {
	contract, leaves := loadCoverageInputs(t)
	for i := range contract.Entries {
		if contract.Entries[i].ConsumedBy != "" {
			contract.Entries[i].ConsumedBy = "not.present"
			break
		}
	}
	if err := contract.Audit(leaves); err == nil || !strings.Contains(err.Error(), "invalid=") {
		t.Fatalf("got %v", err)
	}
}
