package middlewares

import (
	"errors"
	"goadmin/pkg/result"
	"goadmin/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth 用户登录验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		identity := session.GetIdentity(c)

		if identity.ID == 0 {
			if isXhr(c) {
				result.ErrAuth(result.Err(errors.New("登录已过期，刷新浏览器重新登录"))).JSON(c)
			} else {
				c.Redirect(http.StatusFound, "/login")
			}

			c.Abort()

			return
		}

		c.Next()
	}
}
