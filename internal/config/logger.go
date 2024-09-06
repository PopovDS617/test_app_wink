package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	stageEnvName = "STAGE"
)

type loggerConfig struct {
	stage string
}

func NewLoggerConfig() (LoggerConfig, error) {

	stage := os.Getenv(stageEnvName)
	if len(stage) == 0 {
		return nil, errors.New("stage not found")
	}

	return &loggerConfig{
		stage,
	}, nil
}

func (cfg *loggerConfig) Stage() string {
	return cfg.stage
}
