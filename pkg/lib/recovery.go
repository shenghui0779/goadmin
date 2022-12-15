package lib

import (
	"context"
	"goadmin/pkg/logger"
	"runtime/debug"

	"go.uber.org/zap"
)

// Recover recover panic
func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		logger.Err(ctx, "Server Panic",
			zap.Any("error", err),
			zap.ByteString("stack", debug.Stack()),
		)
	}
}

// CtxCopyWithReqID returns a new context with request_id from origin context.
func CtxCopyWithReqID(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), logger.RequestIDKey, logger.GetReqID(ctx))
}
