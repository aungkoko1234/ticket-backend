package service

import (
	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UsersService service.UsersService
	Validate  *validator.Validate
}

func NewAuthServiceImpl (userService  service.UsersService, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UsersService: userService,
		Validate: validate,
	}
}

func (a AuthServiceImpl) Register (register request.RegisterRequest) {
	err := a.Validate.Struct(register)

	helper.ErrorPanic(err)
	a.UsersService.Create(request.CreateUserRequest(register))

}

func (a AuthServiceImpl) Login (login request.LoginRequest) {
	err := a.Validate.Struct(login)

	helper.ErrorPanic(err)
}