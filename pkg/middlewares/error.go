package middlewares

import (
	"errors"
	"goadmin/pkg/logger"
	"goadmin/pkg/result"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Error 处理400和500请求以及panic捕获
func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Err(c.Request.Context(), "Server Panic",
					zap.Any("error", err),
					zap.ByteString("stack", debug.Stack()),
				)

				if isXhr(c) {
					result.ErrSystem().JSON(c)
				} else {
					c.Redirect(http.StatusFound, "/500")
				}

				c.Abort()

				return
			}
		}()

		switch c.Writer.Status() {
		case http.StatusNotFound:
			if isXhr(c) {
				result.ErrNotFound(result.Err(errors.New("page not found")))
			} else {
				c.Redirect(http.StatusFound, "/404")
			}

			c.Abort()

			return
		case http.StatusInternalServerError:
			if isXhr(c) {
				result.ErrSystem().JSON(c)
			} else {
				c.Redirect(http.StatusFound, "/500")
			}

			c.Abort()

			return
		}

		c.Next()
	}
}

// isXhr checks if a request is xml-http-request (ajax).
func isXhr(c *gin.Context) bool {
	x := c.Request.Header.Get("X-Requested-With")

	if strings.ToLower(x) == "xmlhttprequest" {
		return true
	}

	return false
}
