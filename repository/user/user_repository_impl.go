package repository

import (
	"errors"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	"github.com/aungkoko1234/tickermaster_backend/model"
	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl (Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (u UsersRepositoryImpl) Save (user model.Users) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (u UsersRepositoryImpl) Update (user model.Users) {
	var updateUser = request.UpdateUserRequest{Id : user.Id,Name: user.Name,Email: user.Email}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

func (u UsersRepositoryImpl) Delete (userId  int) {
	var user model.Users
	result := u.Db.Where("id = ?",userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (u UsersRepositoryImpl) FindById (userId int) (model.Users,error) {
	var user model.Users
	result := u.Db.Find(&user,userId)

	if result != nil {
		return user,nil
	}else {
		return user, errors.New("user is not found")
	}

}

func (u UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	results := u.Db.Find(&users)
	
	helper.ErrorPanic(results.Error)
	return users
}

func (u UsersRepositoryImpl) FindByEmail(email string) (model.Users,error) {
	var user model.Users
	result := u.Db.Where("email = ?", email).Take(&user)

	if result != nil {
		return user,nil
	}else {
		return user, errors.New("user is not found")
	}
}

