package main

import (
	"flag"
	"fmt"
	"goadmin/pkg/config"
	"goadmin/pkg/ent"
	"goadmin/pkg/html"
	"goadmin/pkg/middlewares"
	"goadmin/pkg/routes"
	"goadmin/pkg/session"
	"net/http"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var envFile string

func main() {
	flag.StringVar(&envFile, "envfile", ".env", "设置ENV配置文件")

	flag.Parse()

	yiigo.LoadEnv(yiigo.WithEnvFile(envFile), yiigo.WithEnvWatcher(func(e fsnotify.Event) {
		yiigo.Logger().Info("env change ok", zap.String("event", e.String()))
		config.RefreshENV()
	}))

	yiigo.Init(
		yiigo.WithMySQL(yiigo.Default, config.DB()),
		yiigo.WithLogger(yiigo.Default, config.Logger()),
	)

	config.RefreshENV()
	ent.InitDB()

	// make sure we have a working tempdir in minimal containers, because:
	// os.TempDir(): The directory is neither guaranteed to exist nor have accessible permissions.
	if err := os.MkdirAll(os.TempDir(), 0775); err != nil {
		yiigo.Logger().Error("err create temp dir", zap.Error(err))
	}

	session.Start()

	serving()
}

func serving() {
	// 弃用Gin内置验证器
	binding.Validator = yiigo.NewValidator()

	r := gin.New()

	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFS("/assets", rice.MustFindBox("./assets").HTTPBox())

	r.Use(middlewares.RequestID(), middlewares.Error(), middlewares.Logger())

	r.HTMLRender = html.NewRender(rice.MustFindBox("./html"))

	routes.Register(r)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	fmt.Println("listening on", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		yiigo.Logger().Fatal("serving error", zap.Error(err))
	}
}
