package serverstatusv5

import (
	"io/ioutil"
	"testing"
)

func loadPolicyContract(t *testing.T) *Contract {
	t.Helper()
	data, err := ioutil.ReadFile("contract.json")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := ParseContract(data)
	if err != nil {
		t.Fatal(err)
	}
	return contract
}

func TestContractMappingsOmissionsVocabulariesAndCeilings(t *testing.T) {
	if err := loadPolicyContract(t).ValidatePolicy(); err != nil {
		t.Fatal(err)
	}
}

func TestContractCeilingTestDetectsLowerLimit(t *testing.T) {
	contract := loadPolicyContract(t)
	contract.SeriesCeilings.Union--
	if err := contract.ValidatePolicy(); err == nil {
		t.Fatal("lower union ceiling was accepted")
	}
}

func TestContractCeilingRejectsAdditionalSeries(t *testing.T) {
	contract := loadPolicyContract(t)
	extra := contract.Entries[0]
	extra.Path = "extra.series"
	extra.PathSegments = []string{"extra", "series"}
	extra.Metric = &MetricContract{Family: "mongodb_server_status_extra_series", Type: "gauge", Unit: "count", Conversion: 1, Help: "extra", Labels: map[string]string{}, SeriesCeiling: 1}
	extra.DecisionSource = "test"
	contract.Entries = append(contract.Entries, extra)
	if err := contract.ValidatePolicy(); err == nil {
		t.Fatal("additional family was accepted without revised ceilings")
	}
}
