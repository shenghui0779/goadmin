package result

import (
	"goadmin/pkg/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	version  = "1.0.0"
	errTitle = "GoAdmin | 错误"
)

// Render render view pages
func Render(c *gin.Context, name string, data ...gin.H) {
	obj := gin.H{}

	if len(data) > 0 {
		obj = data[0]
	}

	obj["version"] = version
	obj["identity"] = session.GetIdentity(c)

	c.HTML(http.StatusOK, name, obj)
}

// Abort render an error page
func Abort(c *gin.Context, code int, msg string) {
	c.HTML(http.StatusOK, "error", gin.H{
		"title": errTitle,
		"code":  code,
		"msg":   msg,
	})
}

// Page403 render a forbid page
func Page403(c *gin.Context) {
	Abort(c, 403, "无操作权限")
}

// Page404 render a not found page
func Page404(c *gin.Context) {
	Abort(c, 404, "页面不存在")
}

// Page500 render a server error page
func Page500(c *gin.Context) {
	Abort(c, 500, "服务器错误")
}

// Redirect redirect to new location
func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}
