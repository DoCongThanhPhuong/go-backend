package middlerwares

import (
	"github.com/DoCongThanhPhuong/go-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, 401, "AUTH::INVALID_TOKEN")
			c.Abort()
			return
		}
		c.Next()
	}
}