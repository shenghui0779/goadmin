package controllers

import (
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/models"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

const version = "1.0.0"
const errTitle = "GoAdmin | 错误"

var captchaDriver = base64Captcha.NewDriverString(39, 120, 0, base64Captcha.OptionShowHollowLine, 4, base64Captcha.TxtNumbers+base64Captcha.TxtAlphabet, nil, nil)

var etcdcli = new(Etcd)

func InitEtcd(etcdaddress string) {
	etcdcli = NewEtcd(etcdaddress)
}

// Render render view pages
func Render(c *gin.Context, name string, data ...gin.H) {
	obj := gin.H{}

	if len(data) > 0 {
		obj = data[0]
	}

	obj["version"] = version

	identity, err := Identity(c)

	if err != nil {
		obj["identity"] = new(models.Identity)
	} else {
		obj["identity"] = identity
	}

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

// ErrForbidden render a forbid page
func ErrForbidden(c *gin.Context) {
	Abort(c, 403, "无操作权限")
}

// ErrNotFound render a not found page
func ErrNotFound(c *gin.Context) {
	Abort(c, 404, "页面不存在")
}

// ErrServer render a server error page
func ErrServer(c *gin.Context) {
	Abort(c, 500, "服务器错误")
}

// Redirect redirect to new location
func Redirect(c *gin.Context, location string) {
	c.Redirect(http.StatusFound, location)
}

// OK returns success of an API.
func OK(c *gin.Context, data ...interface{}) {
	obj := gin.H{
		"err":  false,
		"code": 1000,
		"msg":  "success",
	}

	if len(data) > 0 {
		obj["data"] = data[0]
	}

	c.Set("response", obj)

	c.JSON(http.StatusOK, obj)
}

// Err returns error of an API.
func Err(c *gin.Context, err error, msg ...string) {
	obj := gin.H{
		"err":  true,
		"code": 50000,
		"msg":  "服务器错误，请稍后重试",
	}

	if e, ok := err.(helpers.StatusErr); ok {
		obj["code"] = e.Code()
		obj["msg"] = e.Error()
	}

	if len(msg) > 0 {
		obj["msg"] = msg[0]
	}

	c.Set("response", obj)

	c.JSON(http.StatusOK, obj)
}

// Identity returns an identity.
func Identity(c *gin.Context) (*models.Identity, error) {
	v, ok := c.Get("identity")

	if !ok {
		return nil, errors.New("empty identity")
	}

	identity, ok := v.(*models.Identity)

	if !ok {
		return nil, errors.New("invalid identity")
	}

	if identity == nil {
		return nil, errors.New("empty identity")
	}

	return identity, nil
}
