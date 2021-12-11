package middlewares

import (
	"context"
	"fmt"
	"goadmin/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		if reqID, ok := ctx.Value(logger.RequestIDKey).(string); !ok || len(reqID) == 0 {
			ctx = context.WithValue(ctx, logger.RequestIDKey, fmt.Sprintf("%s-%06d", logger.ReqPrefix, logger.NextRequestID()))
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
