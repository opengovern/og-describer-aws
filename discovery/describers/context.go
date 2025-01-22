package describers

import (
	"context"

	"go.uber.org/zap"
)

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, "logger", logger)
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value("logger").(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return logger
}
