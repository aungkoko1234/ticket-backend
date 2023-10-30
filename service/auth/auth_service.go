package auth_service

import (
	"github.com/aungkoko1234/tickermaster_backend/data/request"
)

type AuthService interface {
	Register(register request.RegisterRequest)
	Login(login request.LoginRequest)
}