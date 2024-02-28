package controller

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/aungkoko1234/tickermaster_backend/utils"
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
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		message := utils.ParseError(err)
        utils.ValidationErrorReponse(ctx, http.StatusBadRequest, http.MethodPost, message)
		return
	}

	token,err := controller.usersService.LoginCheck(loginRequest)

	if err != nil {
		utils.ValidationErrorReponse(ctx, http.StatusUnauthorized, http.MethodPost,gin.H{"message": "username or password is incorrect."})
		return
	}
	utils.ApiResponse(ctx, "Post", http.StatusOK, http.MethodPost, gin.H{"token": token})
}


func (controller *AuthController) Register(ctx *gin.Context){
	registerRequest :=request.RegisterRequest{}
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		message := utils.ParseError(err)
        utils.ValidationErrorReponse(ctx, http.StatusBadRequest, http.MethodPost, message)
		return
	}

	controller.usersService.Create(request.CreateUserRequest(registerRequest))

	utils.ApiResponse(ctx, "Register Success", http.StatusOK, http.MethodPost, gin.H{"message": "Your account is successfully registered"})
}