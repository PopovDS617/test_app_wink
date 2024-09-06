package logger

import (
	"log/slog"
	"os"
)

const (
	stageLocal = "local"
	stageDev   = "dev"
	stageProd  = "prod"
)

var logger *slog.Logger

func setupLogger(stage string) *slog.Logger {
	var log *slog.Logger

	switch stage {
	case stageLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case stageDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case stageProd:
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
