package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5"
)

func TestGeneratedDefinitionsAreCurrent(t *testing.T) {
	contractPath := filepath.Join("..", "..", "collector", "serverstatusv5", "contract.json")
	contractData, err := ioutil.ReadFile(contractPath)
	if err != nil {
		t.Fatal(err)
	}
	contract, err := serverstatusv5.ParseContract(contractData)
	if err != nil {
		t.Fatal(err)
	}
	if err := contract.ValidatePolicy(); err != nil {
		t.Fatal(err)
	}
	generated, err := generate(contract)
	if err != nil {
		t.Fatal(err)
	}
	outputPath := filepath.Join("..", "..", "collector", "serverstatusv5", "definitions_generated.go")
	current, err := ioutil.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(current, generated) {
		t.Fatalf("%s is stale; run go run ./cmd/serverstatus-definitions", outputPath)
	}
}
