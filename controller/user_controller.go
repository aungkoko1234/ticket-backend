package controller

import (
	"net/http"
	"strconv"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
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
	createTagRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
    helper.ErrorPanic(err)

	controller.usersService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK,webResponse)
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
    helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
    helper.ErrorPanic(err)

	updateUserRequest.Id = id

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
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)
	controller.usersService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	userResponse := controller.usersService.FindById(id)

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