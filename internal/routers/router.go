package routers

import (
	"github.com/DoCongThanhPhuong/go-backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

  v1 := r.Group("api/v1") 
  {
    v1.GET("/users/:id", controllers.NewUserController().GetUserById)
  }

	return r
}
