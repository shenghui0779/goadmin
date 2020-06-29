package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iiinsomnia/goadmin/assets"
	"github.com/iiinsomnia/goadmin/middlewares"
	"github.com/iiinsomnia/goadmin/routes"
	"github.com/iiinsomnia/goadmin/session"
	"github.com/iiinsomnia/goadmin/views"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

func main() {
	session.Start()

	loadStaticResource()

	run()
}

// load static resource
func loadStaticResource() {
	assets.LoadAssets()
	views.LoadViews()
}

func run() {
	// 弃用Gin内置验证器
	binding.Validator = yiigo.NewGinValidator()

	debug := yiigo.Env("app.debug").Bool(true)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middlewares.Error())

	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFS("/assets", assets.AssetBox.HTTPBox())

	r.HTMLRender = views.NewRender()

	routes.RouteRegister(r)

	// Graceful restart & zero downtime deploy for Go servers.
	// Use `kill -USR2 pid` to restart.
	if err := gracehttp.Serve(&http.Server{
		Addr:         fmt.Sprintf(":%d", yiigo.Env("app.port").Int(8000)),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}); err != nil {
		yiigo.Logger().Panic("goadmin serving error", zap.String("error", err.Error()))
	}
}
