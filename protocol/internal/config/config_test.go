package config

import (
	"testing"
	"time"
)

func TestLoad_Success(t *testing.T) {
	mockEnv := map[string]string{
		"CCLC_ENVIRONMENT":  "production",
		"CCLC_LOG_LEVEL":    "debug",
		"CCLC_RPC_ENDPOINT": "https://eth.llamprace.com",
		"CCLC_RPC_TIMEOUT":  "5s",
	}

	mockGetter := func(key string) string {
		return mockEnv[key]
	}
	cfg, err := Load(mockGetter)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cfg.App.Environment != "production" {
		t.Errorf("expected environment 'production', got '%s'", cfg.App.Environment)
	}
	if cfg.App.LogLevel != "debug" {
		t.Errorf("expected log level 'debug', got '%s'", cfg.App.LogLevel)
	}
	if cfg.RPC.Endpoint != "https://eth.llamprace.com" {
		t.Errorf("expected endpoint : , got '%s'", cfg.RPC.Endpoint)
	}
	if cfg.RPC.Timeout != 5*time.Second {
		t.Errorf("expected timeout 5s, got '%v'", cfg.RPC.Timeout)
	}
}
func TestLoad_MissingEndpoint(t *testing.T) {
	mockGetter := func(key string) string {
		return ""
	}
	_, err := Load(mockGetter)
	if err == nil {
		t.Fatal("expected error due to missing rpc endpoint, got nil")
	}
}

func TestLoad_Defaults(t *testing.T) {
	mockEnv := map[string]string{
		"CCLC_RPC_ENDPOINT": "https://eth.llamprace.com",
	}
	mockGetter := func(key string) string {
		return mockEnv[key]
	}
	cfg, err := Load(mockGetter)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cfg.App.Environment != "development" {
		t.Errorf("expected default environment development, got '%s'", cfg.App.Environment)
	}
	if cfg.RPC.Timeout != 10*time.Second {
		t.Errorf("expected default timeout 10s, got '%v'", cfg.RPC.Timeout)
	}
}
func TestLoad_InvalidTimeout(t *testing.T) {
	mockEnv := map[string]string{
		"CCLC_RPC_ENDPOINT": "https://eth.llamprace.com",
		"CCLC_RPC_TIMEOUT":  "invalid-duration",
	}
	mockGetter := func(key string) string {
		return mockEnv[key]
	}
	_, err := Load(mockGetter)
	if err == nil {
		t.Fatal("expected error due to invalid CCLC_RPC_TIMEOUT, got nil")
	}
}
