package routers

import (
	"github.com/DoCongThanhPhuong/go-backend/internal/controllers"
	middlerwares "github.com/DoCongThanhPhuong/go-backend/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before => AA")
// 		c.Next()
// 		fmt.Println("Alter => AA")
// 	}
// }
// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before => BB")
// 		c.Next()
// 		fmt.Println("Alter => BB")
// 	}
// }
// func CC(c *gin.Context) {
// 	fmt.Println("Before => CC")
// 	c.Next()
// 	fmt.Println("Alter => CC")
// }

func NewRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(AA(), BB(), CC)
	r.Use(middlerwares.AuthenMiddleware())

  v1 := r.Group("api/v1") 
  {
    v1.GET("/users/:id", controllers.NewUserController().GetUserById)
  }

	return r
}
