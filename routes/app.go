package routes

import (
	"github.com/iiinsomnia/goadmin/controllers"
	"github.com/iiinsomnia/goadmin/middlewares"

	"github.com/gin-gonic/gin"
)

// RouteRegister register routes
func RouteRegister(e *gin.Engine) {
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.Login)
	e.GET("/captcha", controllers.Captcha)
	e.GET("/404", controllers.NotFound)
	e.GET("/500", controllers.InternalServerError)

	root := e.Group("/")
	root.Use(middlewares.Auth())
	{
		root.GET("/", controllers.Home)
		root.GET("/weibo/users", controllers.WeiboUsers)
		root.GET("/weibo/events", controllers.WeiboEvents)

		// 微博
		root.POST("/weibo/user/query", controllers.WeiboUsersQuery)
		root.POST("/weibo/user/add", controllers.WeiboUsersAdd)
		root.POST("/weibo/user/edit", controllers.WeiboUsersUpdate)
		root.POST("/weibo/user/delete", controllers.WeiboUsersDelete)

		// 轰炸
		root.GET("/attack/email", controllers.AttackEmail)
		root.POST("/attack/email/query", controllers.AttackEmailQuery)
		root.POST("/attack/email/again", controllers.AttackEmailAgain)

		root.GET("/logout", controllers.Logout)
		// user
		root.GET("/users", controllers.UserIndex)
		// password
		root.GET("/password/change", controllers.PasswordChange)

		logger := root.Group("/")
		logger.Use(middlewares.Logger())
		{
			// user
			logger.POST("/users/query", controllers.UserQuery)
			logger.POST("/users/add", controllers.UserAdd)
			logger.POST("/users/edit", controllers.UserEdit)
			logger.POST("/users/delete", controllers.UserDelete)
			// password
			logger.POST("/password/change", controllers.PasswordChange)
			logger.POST("/password/reset", controllers.PasswordReset)
		}
	}
}
