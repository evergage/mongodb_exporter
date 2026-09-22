package serverstatusv5

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestModernDescriptorsDoNotCollideWithLegacyInventory(t *testing.T) {
	data, err := ioutil.ReadFile("legacy_descriptors.json")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	if err := json.Unmarshal(data, &names); err != nil {
		t.Fatal(err)
	}
	legacy := make(map[string]bool, len(names))
	for _, name := range names {
		if legacy[name] {
			t.Fatalf("duplicate legacy descriptor %s", name)
		}
		legacy[name] = true
	}
	contractData, err := ioutil.ReadFile("contract.json")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := ParseContract(contractData)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range contract.Entries {
		if entry.Metric != nil && legacy[entry.Metric.Family] {
			t.Errorf("modern descriptor %s for %s collides with legacy inventory", entry.Metric.Family, entry.Path)
		}
	}
}
