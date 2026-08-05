package config

import (
	"strings"
	"testing"
	"time"
)

func TestLoadRejectsMissingDatabaseURL(t *testing.T) {
	t.Parallel()

	_, err := Load(func(string) (string, bool) { return "", false })
	if err == nil || !strings.Contains(err.Error(), "DATABASE_URL") {
		t.Fatalf("expected missing DATABASE_URL error, got %v", err)
	}
}

func TestLoadRejectsMissingRedisURL(t *testing.T) {
	t.Parallel()

	values := map[string]string{"DATABASE_URL": "postgres://localhost/indexer"}
	_, err := Load(mapLookup(values))
	if err == nil || !strings.Contains(err.Error(), "REDIS_URL") {
		t.Fatalf("expected missing REDIS_URL error, got %v", err)
	}
}

func TestLoadRejectsMissingManifestPath(t *testing.T) {
	t.Parallel()

	values := map[string]string{
		"DATABASE_URL": "postgres://localhost/indexer",
		"REDIS_URL":    "redis://localhost:6379/0",
	}
	_, err := Load(mapLookup(values))
	if err == nil || !strings.Contains(err.Error(), "MANIFEST_PATH") {
		t.Fatalf("expected missing MANIFEST_PATH error, got %v", err)
	}
}

func TestLoadUsesSafeDefaults(t *testing.T) {
	t.Parallel()

	values := requiredValues()
	got, err := Load(mapLookup(values))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if got.MetricsAddress != "127.0.0.1:9090" {
		t.Errorf("MetricsAddress = %q", got.MetricsAddress)
	}
	if got.PollInterval != time.Second {
		t.Errorf("PollInterval = %s", got.PollInterval)
	}
	if got.RequestTimeout != 10*time.Second {
		t.Errorf("RequestTimeout = %s", got.RequestTimeout)
	}
	if got.QueueCapacity != 256 {
		t.Errorf("QueueCapacity = %d", got.QueueCapacity)
	}
	if got.ShutdownTimeout != 15*time.Second {
		t.Errorf("ShutdownTimeout = %s", got.ShutdownTimeout)
	}
}

func TestLoadAppliesValidOverrides(t *testing.T) {
	t.Parallel()

	values := requiredValues()
	values["METRICS_ADDRESS"] = "127.0.0.1:9191"
	values["POLL_INTERVAL"] = "2s"
	values["RPC_REQUEST_TIMEOUT"] = "7s"
	values["QUEUE_CAPACITY"] = "64"
	values["SHUTDOWN_TIMEOUT"] = "20s"

	got, err := Load(mapLookup(values))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if got.MetricsAddress != "127.0.0.1:9191" {
		t.Errorf("MetricsAddress = %q", got.MetricsAddress)
	}
	if got.PollInterval != 2*time.Second {
		t.Errorf("PollInterval = %s", got.PollInterval)
	}
	if got.RequestTimeout != 7*time.Second {
		t.Errorf("RequestTimeout = %s", got.RequestTimeout)
	}
	if got.QueueCapacity != 64 {
		t.Errorf("QueueCapacity = %d", got.QueueCapacity)
	}
	if got.ShutdownTimeout != 20*time.Second {
		t.Errorf("ShutdownTimeout = %s", got.ShutdownTimeout)
	}
}

func TestLoadRejectsUnsafeOverrides(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{name: "empty metrics address", key: "METRICS_ADDRESS", value: ""},
		{name: "malformed metrics address", key: "METRICS_ADDRESS", value: "localhost"},
		{name: "zero poll interval", key: "POLL_INTERVAL", value: "0s"},
		{name: "excessive poll interval", key: "POLL_INTERVAL", value: "61s"},
		{name: "negative request timeout", key: "RPC_REQUEST_TIMEOUT", value: "-1s"},
		{name: "excessive request timeout", key: "RPC_REQUEST_TIMEOUT", value: "121s"},
		{name: "zero queue capacity", key: "QUEUE_CAPACITY", value: "0"},
		{name: "excessive queue capacity", key: "QUEUE_CAPACITY", value: "65537"},
		{name: "zero shutdown timeout", key: "SHUTDOWN_TIMEOUT", value: "0s"},
		{name: "excessive shutdown timeout", key: "SHUTDOWN_TIMEOUT", value: "301s"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			values := requiredValues()
			values[test.key] = test.value
			_, err := Load(mapLookup(values))
			if err == nil || !strings.Contains(err.Error(), test.key) {
				t.Fatalf("expected %s validation error, got %v", test.key, err)
			}
		})
	}
}

func requiredValues() map[string]string {
	return map[string]string{
		"DATABASE_URL":  "postgres://localhost/indexer",
		"REDIS_URL":     "redis://localhost:6379/0",
		"MANIFEST_PATH": "deployments/examples/local.json",
	}
}

func mapLookup(values map[string]string) LookupEnv {
	return func(key string) (string, bool) {
		value, ok := values[key]
		return value, ok
	}
}
