package controllers

import (
	"github.com/iiinsomnia/goadmin/helpers"
	"github.com/iiinsomnia/goadmin/service"
	"github.com/iiinsomnia/goadmin/session"

	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// Home ...
func Home(c *gin.Context) {
	Render(c, "home", gin.H{
		"menu": "0",
		"os":   runtime.GOOS,
		"cpu":  runtime.NumCPU(),
		"arch": runtime.GOARCH,
		"go":   runtime.Version(),
	})
}

// Login ...
func Login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		if _, ok := c.Get("identity"); ok {
			Redirect(c, "/")

			return
		}

		Render(c, "login", gin.H{
			"title": "GoAdmin | 登录",
		})

		return
	}

	s := new(service.Login)

	if err := c.ShouldBindJSON(s); err != nil {
		Err(c, helpers.Error(helpers.ErrParams, err))

		return
	}

	if err := s.Do(c); err != nil {
		Err(c, err)

		return
	}

	OK(c)
}

// Logout ...
func Logout(c *gin.Context) {
	if err := session.Destroy(c); err != nil {
		yiigo.Logger().Error("logout error", zap.Error(err))
	}

	Redirect(c, "/login")
}

// Captcha ...
func Captcha(c *gin.Context) {
	captcha := base64Captcha.NewCaptcha(captchaDriver, base64Captcha.DefaultMemStore)

	id, content, err := captcha.Generate()

	if err != nil {
		yiigo.Logger().Error("generate captcha error", zap.Error(err))
	}

	data := yiigo.X{
		"id":      id,
		"captcha": content,
	}

	OK(c, data)
}

func Forbidden(c *gin.Context) {
	ErrForbidden(c)
}

func NotFound(c *gin.Context) {
	ErrNotFound(c)
}

func InternalServerError(c *gin.Context) {
	ErrServer(c)
}
