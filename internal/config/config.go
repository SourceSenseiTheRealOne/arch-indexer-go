// Package config loads and validates process configuration without exposing
// environment access to the rest of the indexer.
package config

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// LookupEnv resolves one configuration value by name.
type LookupEnv func(string) (string, bool)

// Config contains process configuration that is not sourced from a release
// manifest.
type Config struct {
	DatabaseURL     string
	RedisURL        string
	ManifestPath    string
	MetricsAddress  string
	PollInterval    time.Duration
	RequestTimeout  time.Duration
	QueueCapacity   int
	ShutdownTimeout time.Duration
}

// Load reads and validates process configuration.
func Load(lookup LookupEnv) (Config, error) {
	databaseURL, ok := lookup("DATABASE_URL")
	if !ok || databaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	redisURL, ok := lookup("REDIS_URL")
	if !ok || redisURL == "" {
		return Config{}, fmt.Errorf("REDIS_URL is required")
	}
	manifestPath, ok := lookup("MANIFEST_PATH")
	if !ok || manifestPath == "" {
		return Config{}, fmt.Errorf("MANIFEST_PATH is required")
	}

	metricsAddress := stringValue(lookup, "METRICS_ADDRESS", "127.0.0.1:9090")
	if _, _, err := net.SplitHostPort(metricsAddress); err != nil {
		return Config{}, fmt.Errorf("METRICS_ADDRESS must be a host:port address: %w", err)
	}
	pollInterval, err := durationValue(lookup, "POLL_INTERVAL", time.Second, time.Minute)
	if err != nil {
		return Config{}, err
	}
	requestTimeout, err := durationValue(lookup, "RPC_REQUEST_TIMEOUT", 10*time.Second, 2*time.Minute)
	if err != nil {
		return Config{}, err
	}
	queueCapacity, err := intValue(lookup, "QUEUE_CAPACITY", 256, 65_536)
	if err != nil {
		return Config{}, err
	}
	shutdownTimeout, err := durationValue(lookup, "SHUTDOWN_TIMEOUT", 15*time.Second, 5*time.Minute)
	if err != nil {
		return Config{}, err
	}

	return Config{
		DatabaseURL:     databaseURL,
		RedisURL:        redisURL,
		ManifestPath:    manifestPath,
		MetricsAddress:  metricsAddress,
		PollInterval:    pollInterval,
		RequestTimeout:  requestTimeout,
		QueueCapacity:   queueCapacity,
		ShutdownTimeout: shutdownTimeout,
	}, nil
}

func stringValue(lookup LookupEnv, key, fallback string) string {
	if value, ok := lookup(key); ok {
		return value
	}
	return fallback
}

func durationValue(
	lookup LookupEnv,
	key string,
	fallback time.Duration,
	maximum time.Duration,
) (time.Duration, error) {
	value, ok := lookup(key)
	if !ok {
		return fallback, nil
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a duration: %w", key, err)
	}
	if parsed <= 0 || parsed > maximum {
		return 0, fmt.Errorf("%s must be greater than zero and at most %s", key, maximum)
	}
	return parsed, nil
}

func intValue(lookup LookupEnv, key string, fallback, maximum int) (int, error) {
	value, ok := lookup(key)
	if !ok {
		return fallback, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	if parsed <= 0 || parsed > maximum {
		return 0, fmt.Errorf("%s must be greater than zero and at most %d", key, maximum)
	}
	return parsed, nil
}
