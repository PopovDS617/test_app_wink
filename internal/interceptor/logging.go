package interceptor

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func LoggingOptions() []logging.Option {
	return []logging.Option{
		logging.WithDisableLoggingFields(logging.SystemTag[0], logging.SystemTag[1], logging.ComponentFieldKey, logging.ServiceFieldKey, logging.MethodFieldKey, logging.MethodTypeFieldKey, logging.KindClientFieldValue, logging.KindServerFieldValue),
		logging.WithLogOnEvents(logging.FinishCall),
	}
}
