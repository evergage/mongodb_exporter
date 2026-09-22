package fixturetool

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGeneratedFixturesAreCurrent(t *testing.T) {
	root := filepath.Join("..", "testdata")
	if err := Generate(root, true); err != nil {
		t.Fatal(err)
	}
}

func TestReconstructionIsDeterministic(t *testing.T) {
	source, err := ioutil.ReadFile(filepath.Join("..", "testdata", "source", "server-status-primary-5.0.34.json"))
	if err != nil {
		t.Fatal(err)
	}
	first, err := ReconstructJSON(source)
	if err != nil {
		t.Fatal(err)
	}
	second, err := ReconstructJSON(source)
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("reconstruction is not deterministic")
	}
}
