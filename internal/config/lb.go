package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	originalHostEnvName = "ORIGINAL_HOST"
	cdnHostEnvName      = "CDN_HOST"
)

type lbConfig struct {
	originalHost string
	cdnHost      string
}

func NewLBConfig() (LoadBalancerConfig, error) {

	originalHost := os.Getenv(originalHostEnvName)
	if len(originalHost) == 0 {
		return nil, errors.New("original host not found")
	}

	cdnHost := os.Getenv(cdnHostEnvName)
	if len(cdnHost) == 0 {
		return nil, errors.New("cdn host not found")
	}

	return &lbConfig{
		originalHost, cdnHost,
	}, nil
}

func (cfg *lbConfig) OriginalHost() string {
	return cfg.originalHost
}

func (cfg *lbConfig) CDNHost() string {
	return cfg.cdnHost
}
