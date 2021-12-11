package routes

import (
	"goadmin/pkg/middlewares"
	"goadmin/pkg/result"
	"goadmin/pkg/service"
	"runtime"

	"github.com/gin-gonic/gin"
)

// Register register routes
func Register(router *gin.Engine) {
	router.GET("/", middlewares.Auth(), func(c *gin.Context) {
		result.Render(c, "home", gin.H{
			"menu": "0",
			"os":   runtime.GOOS,
			"cpu":  runtime.NumCPU(),
			"arch": runtime.GOARCH,
			"go":   runtime.Version(),
		})
	})

	router.GET("/403", func(c *gin.Context) {
		result.Page403(c)
	})

	router.GET("/404", func(c *gin.Context) {
		result.Page404(c)
	})

	router.GET("/500", func(c *gin.Context) {
		result.Page500(c)
	})

	auth(router)
	password(router)
	user(router)
}

func auth(router *gin.Engine) {
	auth := service.NewAuth()

	router.GET("/login", auth.Index)
	router.GET("/captcha", auth.Captcha)
	router.POST("/login", auth.Login)
	router.GET("/logout", middlewares.Auth(), auth.Logout)
}

func password(router *gin.Engine) {
	password := service.NewPassword()

	auth := router.Group("/password")
	auth.Use(middlewares.Auth())
	{
		auth.GET("/", password.Index)
		auth.POST("/change", password.Change)
		auth.GET("/reset/:uid", password.Reset)
	}
}

func user(router *gin.Engine) {
	user := service.NewUser()

	auth := router.Group("/users")
	auth.Use(middlewares.Auth())
	{
		auth.GET("/", user.Index)
		auth.POST("/list", user.List)
		auth.POST("/create", user.Create)
		auth.POST("/update/:uid", user.Update)
		auth.GET("/delete/:uid", user.Delete)
	}
}
