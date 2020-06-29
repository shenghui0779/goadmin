package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/shenghui0779/yiigo"

	"github.com/gin-gonic/gin"

	"github.com/iiinsomnia/goadmin/helpers"
)

// Error 处理400和500请求以及panic捕获
func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				yiigo.Logger().Error(fmt.Sprintf("%v\n%s", err, string(debug.Stack())))

				if isXhr(c) {
					c.JSON(http.StatusOK, gin.H{
						"code":    helpers.ErrSystem,
						"success": false,
						"msg":     "服务器错误",
					})
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
				c.JSON(http.StatusOK, gin.H{
					"code":    helpers.ErrPageNotFound,
					"success": false,
					"msg":     "页面不存在",
				})
			} else {
				c.Redirect(http.StatusFound, "/404")
			}

			c.Abort()

			return
		case http.StatusInternalServerError:
			if isXhr(c) {
				c.JSON(http.StatusOK, gin.H{
					"code":    helpers.ErrSystem,
					"success": false,
					"msg":     "服务器错误",
				})
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
