package fixturetool

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	redactionVersion           = "1"
	reconstructionVersion      = "1"
	normalizationVersion       = "1"
	originalBSONTypesAvailable = false
)

var fixtureMetadata = []struct {
	role, mongoVersion, captureTime, sourceSHA256, reconstructedSHA256 string
}{
	{"primary", "5.0.34", "2026-09-21T21:05:28.899Z", "5c84aec99d1de82306b09f2a44af4e5fd59af38997ff7d0665b2b08f78d57a0b", "808ee6ef1a5c8aed11d41508fe132651e489fccfdbb75aa62d7d7f3e2c83b166"},
	{"secondary", "5.0.34", "2026-09-21T20:42:39.325Z", "227abf474edc8b83ab63311b5c19cecab9e11918c3a819237eadb5875efe8ac7", "f7985c6ed450f694db0bc19603e74770707c06f25eeeade98db3a98a4c844152"},
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

func TestReconstructionRejectsIdentities(t *testing.T) {
	if _, err := ReconstructJSON([]byte("{\"host\":\"prod4.evergage.com\"}")); err == nil {
		t.Fatal("identity-bearing fixture accepted")
	}
}
