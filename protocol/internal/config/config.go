package config

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// hold primary configuration
type config struct {
	App Appconfig
	RPC RPCConfig
}
type Appconfig struct {
	Environment string
	LogLevel    string
}
type RPCConfig struct {
	Endpoint string
	Timeout  time.Duration
}

// inject mock getters during test
type EnvGetter func(string) string

func Load(getEnv EnvGetter) (*config, error) {
	if getEnv == nil {
		getEnv = os.Getenv
	}
	//initialize with sensible defaults
	cfg := &config{
		App: Appconfig{
			Environment: "development",
			LogLevel:    "info",
		},
		RPC: RPCConfig{
			Timeout: 10 * time.Second,
		},
	}
	//load the App configuration
	if env := getEnv("CCLC_ENVIRONMENT"); env != "" {
		cfg.App.Environment = env
	}
	if level := getEnv("CCLC_LOG_LEVEL"); level != "" {
		cfg.App.LogLevel = level
	}
	//Load RPC configuration
	endpoint := getEnv("CCLC_RPC_ENDPOINT")
	if endpoint == "" {
		return nil, errors.New("missing required configuration: CCLC_RPC_ENDPOINT")
	}
	cfg.RPC.Endpoint = endpoint
	//otional timeout
	if timeoutStr := getEnv("CCLC_RPC_TIMEOUT"); timeoutStr != "" {
		timeout, err := time.ParseDuration(timeoutStr)
		if err != nil {
			return nil, fmt.Errorf("invalid CCLC_RPC_TIMEOUT value: %w", err)
		}
		cfg.RPC.Timeout = timeout
	}
	return cfg, nil
}
