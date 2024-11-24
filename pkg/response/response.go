package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	StatusCode    int         `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, code int, message string, data interface{}) {
  ctx.JSON(http.StatusOK, ResponseData{
		StatusCode: code,
		Message: message,
		Data: data,
	})
}

func ErrorResponse(ctx *gin.Context, code int, message string) {
  ctx.JSON(http.StatusOK, ResponseData{
		StatusCode: code,
		Message: message,
    Data: nil,
	})
}