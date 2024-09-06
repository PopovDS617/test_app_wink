package logger

import (
	"log/slog"
	"os"
)

const (
	stageLocalEnvName = "local"
	stageDevEnvName   = "dev"
	stageProdEnvName  = "prod"
)

var logger *slog.Logger

func setupLogger(stage string) *slog.Logger {
	var log *slog.Logger

	switch stage {
	case stageLocalEnvName:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case stageDevEnvName:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case stageProdEnvName:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func NewLogger(stage string) *slog.Logger {
	logger = setupLogger(stage)
	return logger
}
