package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	cdnHostEnvName = "CDN_HOST"
)

type lbConfig struct {
	cdnHost string
}

func NewLBConfig() (LoadBalancerConfig, error) {

	cdnHost := os.Getenv(cdnHostEnvName)
	if len(cdnHost) == 0 {
		return nil, errors.New("cdn host not found")
	}

	return &lbConfig{
		cdnHost,
	}, nil
}

func (cfg *lbConfig) CDNHost() string {
	return cfg.cdnHost
}
