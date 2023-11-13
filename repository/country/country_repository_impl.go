package country_repository

import (
	"errors"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	"github.com/aungkoko1234/tickermaster_backend/model"
	"gorm.io/gorm"
)

type CountriesRepositoryImpl struct {
	Db *gorm.DB
}

func NewCountriesRepositoryImpl(Db * gorm.DB) CountriesRepository {
	return  &CountriesRepositoryImpl{Db: Db}
}

func (c CountriesRepositoryImpl) Save(country model.Countries) {
	result := c.Db.Create(&country)
	helper.ErrorPanic(result.Error)
}

func (c CountriesRepositoryImpl) Update (country model.Countries) {
	var updateCountry = request.UpdateCountryRequest{Id : country.ID,  Name: country.Name,Code: country.Code}
	result := c.Db.Model(&country).Updates(updateCountry)
	helper.ErrorPanic(result.Error)
}

func (c CountriesRepositoryImpl) Delete (countryId  string) {
	var user model.Users
	result := c.Db.Where("id = ?",countryId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

func (c CountriesRepositoryImpl) FindById (countryId string) (model.Countries,error) {
	var country model.Countries
	result := c.Db.Find(&country,countryId)

	if result != nil {
		return country,nil
	}else {
		return country, errors.New("country is not found")
	}

}

func (c CountriesRepositoryImpl) FindAll() []model.Countries {
	var countries []model.Countries
	results := c.Db.Find(&countries)
	
	helper.ErrorPanic(results.Error)
	return countries
}