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

func TestExpandedCommandContract(t *testing.T) {
	contract := loadPolicyContract(t)
	familyCounts := map[string]int{}
	commandNames := map[string]bool{}
	validatorKinds := map[string]bool{}
	unknownCommands := 0
	for _, entry := range contract.Entries {
		if entry.Metric == nil || len(entry.PathSegments) < 3 || entry.PathSegments[0] != "metrics" || entry.PathSegments[1] != "commands" {
			continue
		}
		familyCounts[entry.Metric.Family]++
		switch entry.Metric.Family {
		case "mongodb_server_status_commands_total":
			commandNames[entry.Metric.Labels["command"]] = true
			if entry.Metric.Help != "Total MongoDB serverStatus command executions by command and outcome; total includes failed executions." {
				t.Errorf("%s has leaf-specific HELP %q", entry.Path, entry.Metric.Help)
			}
		case "mongodb_server_status_command_validator_total":
			validatorKinds[entry.Metric.Labels["command"]+"\x00"+entry.Metric.Labels["kind"]] = true
		case "mongodb_server_status_unknown_commands_total":
			unknownCommands++
			if len(entry.Metric.Labels) != 0 {
				t.Errorf("unknown command sentinel has labels %v", entry.Metric.Labels)
			}
		}
	}
	if len(commandNames) != 218 || familyCounts["mongodb_server_status_commands_total"] != 436 {
		t.Fatalf("command names=%d command outcomes=%d, want 218 and 436", len(commandNames), familyCounts["mongodb_server_status_commands_total"])
	}
	if len(validatorKinds) != 6 || familyCounts["mongodb_server_status_command_validator_total"] != 6 {
		t.Fatalf("validator kinds=%d samples=%d, want 6", len(validatorKinds), familyCounts["mongodb_server_status_command_validator_total"])
	}
	if unknownCommands != 1 {
		t.Fatalf("unknown command sentinel mappings=%d, want 1", unknownCommands)
	}
	if contract.SeriesCeilings.PerFamily != 436 {
		t.Fatalf("per-family ceiling=%d, want 436", contract.SeriesCeilings.PerFamily)
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
