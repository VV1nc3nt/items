package logger

import (
	"log/slog"
	"os"
)

const envLogLevel = "LOG_LEVEL"

func New() *slog.Logger {
	logLevel, _ := os.LookupEnv(envLogLevel)
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     parseLogLevel(logLevel),
	})
	return slog.New(handler)
}

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}