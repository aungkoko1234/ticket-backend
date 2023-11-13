package country_repository

import "github.com/aungkoko1234/tickermaster_backend/model"

type CountriesRepository interface {
	Save (user model.Countries)
	Update (user model.Countries)
	Delete (userId string)
	FindById (userId string) (user model.Countries,err error)
	FindAll () []model.Countries
}