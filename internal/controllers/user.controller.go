package controllers

import (
	"net/http"

	"github.com/DoCongThanhPhuong/go-backend/internal/services"
	"github.com/DoCongThanhPhuong/go-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	response.SuccessResponse(ctx, http.StatusOK, "USER::GET_DETAILS_SUCCESS", uc.userService.GetUserByID(id))
}