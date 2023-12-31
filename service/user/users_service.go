package user_service

import (
	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
)

type UsersService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete (userId string)
	FindById(userId string) response.UserResponse
	FindAll() []response.UserResponse
	LoginCheck(login request.LoginRequest) (string,error)
}