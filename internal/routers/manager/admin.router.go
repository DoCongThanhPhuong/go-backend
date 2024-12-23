package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup){
	adminRouter := Router.Group("/admin/user")
	{
		adminRouter.GET("/active_user")
	}
}