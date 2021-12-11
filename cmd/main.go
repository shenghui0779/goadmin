package main

import (
	"fmt"
	"goadmin/pkg/config"
	"goadmin/pkg/console"
	"goadmin/pkg/ent"
	"goadmin/pkg/html"
	"goadmin/pkg/middlewares"
	"goadmin/pkg/routes"
	"goadmin/pkg/session"
	"net/http"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenghui0779/yiigo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var envFile string

func main() {
	app := &cli.App{
		Name:     "goadmin",
		Usage:    "go web project template",
		Commands: console.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "envfile",
				Aliases:     []string{"E"},
				Value:       ".env",
				Usage:       "设置配置文件，默认：.env",
				Destination: &envFile,
			},
		},
		Before: func(c *cli.Context) error {
			yiigo.LoadEnv(yiigo.WithEnvFile(envFile))

			yiigo.Init(
				yiigo.WithMySQL(yiigo.Default, config.DB()),
				yiigo.WithLogger(yiigo.Default, config.Logger()),
			)

			ent.InitDB()
			session.Start()

			return nil
		},
		Action: func(c *cli.Context) error {
			serving()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		yiigo.Logger().Fatal("app running error", zap.Error(err))
	}
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
