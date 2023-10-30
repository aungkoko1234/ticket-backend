package user_service

import (
	"fmt"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	"github.com/aungkoko1234/tickermaster_backend/model"
	repository "github.com/aungkoko1234/tickermaster_backend/repository/user"
	"github.com/aungkoko1234/tickermaster_backend/utils/token"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UsersServiceImpl struct {
	UserRepository repository.UsersRepository
	Validate  *validator.Validate
}

func NewUsersServiceImpl (userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UserRepository: userRepository,
		Validate: validate,
	}
}


func (u UsersServiceImpl) Create (user request.CreateUserRequest) {
	err := u.Validate.Struct(user)

	helper.ErrorPanic(err)

	userModel := model.Users {
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	u.UserRepository.Save(userModel)
}

func (u UsersServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)

	helper.ErrorPanic(err)

	userData.Name = user.Name
	userData.Email = user.Email

	u.UserRepository.Update(userData)
}

func (u UsersServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u UsersServiceImpl) FindById(userId int ) response.UserResponse{
	userData, err := u.UserRepository.FindById(userId)

	helper.ErrorPanic(err)

	userResponse := response.UserResponse {
		Id   : userData.Id,
		Name : userData.Name,
		Email: userData.Email ,
	}

	return userResponse
}

func (u UsersServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()

	var users  []response.UserResponse

	for _,value := range result {
		user := response.UserResponse{
			Id: value.Id,
			Name: value.Name,
			Email:  value.Email,
		}
		users = append(users, user)
	}
	
	return users

}

func VerifyPassword(hashedPassword,password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u UsersServiceImpl) LoginCheck(login request.LoginRequest) (string,error) {
	userData,err := u.UserRepository.FindByEmail(login.Email)

    helper.ErrorPanic(err)

	fmt.Println("user",userData.Password)
	fmt.Println("request",login.Password)

	err = VerifyPassword(userData.Password,login.Password) 

	if err != nil {
		return "",err
	}
    
	token,err := token.GenerateToken(userData.ID)

	helper.ErrorPanic(err)

	return token, nil
	

}
