package controller

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	usersService  service.UsersService
}

func NewAuthController (service service.UsersService) *AuthController {
	return &AuthController{
		usersService: service,
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
    helper.ErrorPanic(err)

	token,err := controller.usersService.LoginCheck(loginRequest)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect."})
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   token,
	}

	ctx.JSON(http.StatusOK,webResponse)
}