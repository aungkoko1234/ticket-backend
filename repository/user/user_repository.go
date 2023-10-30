package user_repository

import "github.com/aungkoko1234/tickermaster_backend/model"

type UsersRepository interface {
	Save (user model.Users)
	Update (user model.Users)
	Delete (userId int)
	FindById (userId int) (user model.Users,err error)
	FindAll () []model.Users
	FindByEmail (email string) (user model.Users, err error)
}