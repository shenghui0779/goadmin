package helpers

import (
	"context"
	"goadmin/pkg/logger"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
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

func URLParamInt(c *gin.Context, key string) int64 {
	param := c.Param(key)

	if len(param) == 0 {
		return 0
	}

	v, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		logger.Err(c.Request.Context(), "err url param to int64", zap.Error(err), zap.String("param", key))

		return 0
	}

	return v
}

func URLQueryInt(c *gin.Context, key string) int64 {
	query := c.Query(key)

	if len(query) == 0 {
		return 0
	}

	v, err := strconv.ParseInt(query, 10, 64)

	if err != nil {
		logger.Err(c.Request.Context(), "err url query to int64", zap.Error(err), zap.String("query", key))

		return 0
	}

	return v
}

func Identity() {

}
