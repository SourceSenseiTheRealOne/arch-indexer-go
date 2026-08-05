package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/SourceSenseiTheRealOne/arch-indexer-go/internal/config"
)

const (
	exitOK     = 0
	exitUsage  = 64
	exitConfig = 78
)

func main() {
	os.Exit(run(os.Args[1:], os.LookupEnv, os.Stdout, os.Stderr))
}

func run(args []string, lookup config.LookupEnv, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("indexerd", flag.ContinueOnError)
	flags.SetOutput(stderr)
	checkConfig := flags.Bool("check-config", false, "validate configuration and exit")
	if err := flags.Parse(args); err != nil {
		return exitUsage
	}
	if flags.NArg() != 0 {
		fmt.Fprintln(stderr, "indexerd does not accept positional arguments")
		return exitUsage
	}
	if !*checkConfig {
		fmt.Fprintln(stderr, "indexer runtime is not implemented; use --check-config to validate the foundation")
		return exitUsage
	}

	if _, err := config.Load(lookup); err != nil {
		fmt.Fprintf(stderr, "configuration invalid: %v\n", err)
		return exitConfig
	}

	fmt.Fprintln(stdout, "configuration valid; indexer runtime is not started")
	return exitOK
}
