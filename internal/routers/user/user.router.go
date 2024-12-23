package user

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	// public
	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register")
		publicUserRouter.POST("/otp")
	}

	// private
	privateUserRouter := Router.Group("/user")
	// privateUserRouter.Use(limiter())
	// privateUserRouter.Use(Authen())
	// privateUserRouter.Use(Permission())
	{
		privateUserRouter.GET("/info")
	}
}