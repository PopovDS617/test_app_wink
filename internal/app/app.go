package app

import (
	"context"
	"log"
	"log/slog"
	"net"
	"testappwink/internal/interceptor"
	"testappwink/internal/logger"
	gen "testappwink/pkg/lb_v1"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	logger          *slog.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initLogger,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	a.logger = logger.NewLogger(a.serviceProvider.loggerConfig.Stage())
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {

	loggingInterceptor := interceptor.InterceptorLogger(a.logger)
	logginOptions := interceptor.LoggingOptions()

	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(loggingInterceptor, logginOptions...)))

	reflection.Register(a.grpcServer)

	gen.RegisterLBV1Server(a.grpcServer, a.serviceProvider.initLBImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.grpcConfig.Address())

	list, err := net.Listen("tcp", a.serviceProvider.grpcConfig.Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
