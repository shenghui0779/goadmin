package logger

import (
	"context"
	"fmt"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Info(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Warn(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}

func Err(ctx context.Context, msg string, fields ...zap.Field) {
	yiigo.Logger().Error(fmt.Sprintf("[%s] %s", GetReqID(ctx), msg), fields...)
}
