package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup){
	publicUserRouter := Router.Group("/admin")
	{
    publicUserRouter.POST("/login")
  }

	privateUserRouter := Router.Group("/admin/user")
	// privateUserRouter.Use(Authen())
	// privateUserRouter.Use(Permission())
	{
		privateUserRouter.GET("/active_user")
	}
}