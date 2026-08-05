package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCheckConfigDoesNotPrintConnectionStrings(t *testing.T) {
	t.Parallel()

	values := map[string]string{
		"DATABASE_URL":  "postgres://localhost/indexer",
		"REDIS_URL":     "redis://localhost:6379/0",
		"MANIFEST_PATH": "deployments/examples/local.json",
	}
	lookup := func(key string) (string, bool) {
		value, ok := values[key]
		return value, ok
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	exitCode := run([]string{"--check-config"}, lookup, &stdout, &stderr)

	if exitCode != 0 {
		t.Fatalf("exit code = %d, stderr = %q", exitCode, stderr.String())
	}
	combined := stdout.String() + stderr.String()
	for _, forbidden := range []string{"postgres://", "redis://"} {
		if strings.Contains(combined, forbidden) {
			t.Fatalf("output contains sensitive marker %q: %q", forbidden, combined)
		}
	}
	if !strings.Contains(stdout.String(), "configuration valid") {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestRunRefusesToPretendTheRuntimeExists(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run(nil, func(string) (string, bool) { return "", false }, &stdout, &stderr)

	if exitCode != 64 {
		t.Fatalf("exit code = %d, want 64", exitCode)
	}
	if !strings.Contains(stderr.String(), "runtime is not implemented") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestRunReturnsConfigurationExitCode(t *testing.T) {
	t.Parallel()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	exitCode := run(
		[]string{"--check-config"},
		func(string) (string, bool) { return "", false },
		&stdout,
		&stderr,
	)

	if exitCode != 78 {
		t.Fatalf("exit code = %d, want 78", exitCode)
	}
	if !strings.Contains(stderr.String(), "DATABASE_URL") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}
