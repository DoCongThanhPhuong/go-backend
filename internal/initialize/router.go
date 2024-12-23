package initialize

import (
	"github.com/DoCongThanhPhuong/go-backend/global"
	"github.com/DoCongThanhPhuong/go-backend/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	// r.Use() // logging
	// r.Use() // cross 
	// r.Use() // limiter global
	managerRouter := routers.AppRouterGroup.Manager
	userRouter := routers.AppRouterGroup.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/check-status")
	}
	{
		managerRouter.InitUserRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	return r
}