package logger

import (
	"context"
	"log/slog"
	"testing"
)

func TestInit_DebugLevel(t *testing.T) {
	Init("debug", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Debug level must allow Debug messages")
	}
}

func TestInit_InfoLevel(t *testing.T) {
	Init("info", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelInfo) {
		t.Error("Info level must allow Info messages")
	}
	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Info level must NOT allow Debug messages")
	}
}

func TestInit_WarnLevel(t *testing.T) {
	Init("warn", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelWarn) {
		t.Error("Warn level must allow Warn messages")
	}
	if slog.Default().Enabled(context.Background(), slog.LevelInfo) {
		t.Error("Warn level must NOT allow Info messages")
	}
}

func TestInit_ErrorLevel(t *testing.T) {
	Init("error", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelError) {
		t.Error("Error level must allow Error messages")
	}
	if slog.Default().Enabled(context.Background(), slog.LevelWarn) {
		t.Error("Error level must NOT allow Warn messages")
	}
}

func TestInit_InvalidLevel(t *testing.T) {
	Init("nonsense-word", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelInfo) {
		t.Error("Invalid level must fall back to Info and allow Info messages")
	}
	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Invalid level fallback must NOT allow Debug messages")
	}
}

func TestInit_UppercaseLevel(t *testing.T) {
	Init("DEBUG", "development")

	if !slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Uppercase DEBUG must be recognized and allow Debug messages")
	}
}

func TestInit_ProductionFormat(t *testing.T) {
	Init("info", "production")

	handler := slog.Default().Handler()
	if _, isJSON := handler.(*slog.JSONHandler); !isJSON {
		t.Error("Production environment must format logs as JSON")
	}
}

func TestInit_DevelopmentFormat(t *testing.T) {
	Init("info", "development")

	handler := slog.Default().Handler()
	if _, isText := handler.(*slog.TextHandler); !isText {
		t.Error("Development environment must format logs as plain Text")
	}
}
