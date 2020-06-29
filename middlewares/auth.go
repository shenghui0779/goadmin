package middlewares

import (
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"github.com/iiinsomnia/goadmin/consts"
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/models"
	"github.com/iiinsomnia/goadmin/session"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth 用户登录验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		identity, ok := getIdentity(c)

		if !ok {
			if isXhr(c) {
				c.JSON(http.StatusOK, gin.H{
					"code":    helpers.ErrLoginExpired,
					"success": false,
					"msg":     "登录已过期，刷新浏览器重新登录",
				})
			} else {
				c.Redirect(http.StatusFound, "/login")
			}

			c.Abort()

			return
		}

		c.Set("identity", identity)
		c.Next()
	}
}

func getIdentity(c *gin.Context) (*models.Identity, bool) {
	v, err := session.Get(c, consts.SessionID)

	if err != nil {
		yiigo.Logger().Error("get identity from session error", zap.Error(err))

		return nil, false
	}

	if v == nil {
		return nil, false
	}

	return v.(*models.Identity), true
}
