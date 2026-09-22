package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"

	"github.com/percona/mongodb_exporter/collector/mongod"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/bson"
)

var fqNamePattern = regexp.MustCompile(`fqName: "([^"]+)"`)

func main() {
	output := flag.String("output", "collector/serverstatusv5/legacy_descriptors.json", "inventory path")
	check := flag.Bool("check", false, "verify inventory without changing it")
	flag.Parse()
	names := map[string]bool{}
	for _, path := range []string{
		"collector/serverstatusv5/testdata/reconstructed/server-status-primary-5.0.34.json",
		"collector/serverstatusv5/testdata/reconstructed/server-status-secondary-5.0.34.json",
	} {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fatal(err)
		}
		var raw bson.Raw
		if err := bson.UnmarshalExtJSON(data, true, &raw); err != nil {
			fatal(err)
		}
		var status mongod.ServerStatus
		if err := bson.Unmarshal(raw, &status); err != nil {
			fatal(err)
		}
		metrics := make(chan prometheus.Metric, 10000)
		status.Export(metrics)
		close(metrics)
		for metric := range metrics {
			match := fqNamePattern.FindStringSubmatch(metric.Desc().String())
			if len(match) != 2 {
				fatal(fmt.Errorf("cannot parse descriptor %s", metric.Desc()))
			}
			names[match[1]] = true
		}
	}
	sorted := make([]string, 0, len(names))
	for name := range names {
		sorted = append(sorted, name)
	}
	sort.Strings(sorted)
	data, err := json.MarshalIndent(sorted, "", "  ")
	if err != nil {
		fatal(err)
	}
	data = append(data, '\n')
	if *check {
		current, err := ioutil.ReadFile(*output)
		if err != nil {
			fatal(err)
		}
		if !bytes.Equal(current, data) {
			fatal(fmt.Errorf("%s is stale; run serverstatus-descriptors", *output))
		}
		return
	}
	if err := ioutil.WriteFile(*output, data, 0644); err != nil {
		fatal(err)
	}
}
func fatal(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
