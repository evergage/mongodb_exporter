package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/percona/mongodb_exporter/collector/serverstatusv5/fixturetool"
)

func main() {
	root := flag.String("root", "collector/serverstatusv5/testdata", "fixture directory")
	check := flag.Bool("check", false, "verify generated files without changing them")
	flag.Parse()
	if err := fixturetool.Generate(*root, *check); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
