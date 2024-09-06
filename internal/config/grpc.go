package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return fmt.Sprintf(":%s", cfg.port)
}
