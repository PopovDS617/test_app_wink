package app

import (
	"context"
	"testappwink/internal/api/lb"
	"testappwink/internal/config"
	service "testappwink/internal/service"
	lbService "testappwink/internal/service/lb"
	"testappwink/pkg/utils"
)

type serviceProvider struct {
	grpcConfig   config.GRPCConfig
	lbConfig     config.LoadBalancerConfig
	loggerConfig config.LoggerConfig
	lbService    service.LBService
	lbImpl       *lb.Implementation
}

func newServiceProvider() *serviceProvider {

	srvProvider := serviceProvider{}

	srvProvider.initGRPCConfig()
	srvProvider.initLBConfig()
	srvProvider.initLoggerConfig()

	return &srvProvider
}

func (s *serviceProvider) initLBConfig() {
	if s.lbConfig == nil {
		cfg := utils.Must(config.NewLBConfig())

		s.lbConfig = cfg
	}

}

func (s *serviceProvider) initGRPCConfig() {
	if s.grpcConfig == nil {
		cfg := utils.Must(config.NewGRPCConfig())

		s.grpcConfig = cfg
	}
}

func (s *serviceProvider) initLoggerConfig() {
	if s.loggerConfig == nil {
		cfg := utils.Must(config.NewLoggerConfig())

		s.loggerConfig = cfg
	}

}

func (s *serviceProvider) initLBService(_ context.Context) service.LBService {
	if s.lbService == nil {
		s.lbService = lbService.NewService(s.lbConfig.CDNHost())
	}
	return s.lbService
}

func (s *serviceProvider) initLBImpl(ctx context.Context) *lb.Implementation {
	if s.lbImpl == nil {
		s.lbImpl = lb.NewImplementation(s.initLBService(ctx))
	}

	return s.lbImpl
}
