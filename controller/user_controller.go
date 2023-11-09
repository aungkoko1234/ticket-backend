package controller

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/aungkoko1234/tickermaster_backend/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
  usersService  service.UsersService
}

func NewUserController (service service.UsersService) *UserController {
	return &UserController{
		usersService: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		message := utils.ParseError(err)
        utils.ValidationErrorReponse(ctx, http.StatusBadRequest, http.MethodPost, message)
		return
	}

	controller.usersService.Create(createUserRequest)

	utils.ApiResponse(ctx, "Register Success", http.StatusOK, http.MethodPost, gin.H{"message": "Your account is successfully registered"})
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
    helper.ErrorPanic(err)

	userId := ctx.Param("userId")

	updateUserRequest.Id = userId

	controller.usersService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK,webResponse)
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	// id, err := strconv.Atoi(userId)
	controller.usersService.Delete(userId)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	userResponse := controller.usersService.FindById(userId)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.usersService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}