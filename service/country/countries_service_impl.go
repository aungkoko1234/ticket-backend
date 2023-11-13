package country_service

import (
	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
	"github.com/aungkoko1234/tickermaster_backend/helper"
	"github.com/aungkoko1234/tickermaster_backend/model"
	country_repository "github.com/aungkoko1234/tickermaster_backend/repository/country"
	"github.com/go-playground/validator/v10"
)

type CountriesServiceImpl struct {
	CountryRepository country_repository.CountriesRepository
	Validate  *validator.Validate
}

func NewCountriesServiceImpl (countryRepository  country_repository.CountriesRepository, validate *validator.Validate) CountriesService {
	return &CountriesServiceImpl{
		CountryRepository: countryRepository,
		Validate: validate,
	}
}


func (c CountriesServiceImpl) Create (country request.CreateCountryRequest) {
	err := c.Validate.Struct(country)

	helper.ErrorPanic(err)

	countryModel := model.Countries {
		Name: country.Name,
		Code : country.Code,
	}

	c.CountryRepository.Save(countryModel)
}

func (c CountriesServiceImpl) Update(country request.UpdateCountryRequest) {
	countryData, err := c.CountryRepository.FindById(country.Id)

	helper.ErrorPanic(err)

	countryData.Name = country.Name
	countryData.Code = country.Code

	c.CountryRepository.Update(countryData)
}

func (u CountriesServiceImpl) Delete(userId string) {
	u.CountryRepository.Delete(userId)
}

func (c CountriesServiceImpl) FindById(userId string ) response.CountryResponse{
	countryData, err := c.CountryRepository.FindById(userId)

	helper.ErrorPanic(err)

	CountryResponse := response.CountryResponse {
		Id   : countryData.ID,
		Name : countryData.Name,
		Code : countryData.Code ,
	}

	return CountryResponse
}

func (c CountriesServiceImpl) FindAll() []response.CountryResponse {
	result := c.CountryRepository.FindAll()

	var users  []response.CountryResponse

	for _,value := range result {
		user := response.CountryResponse{
			Id: value.ID,
			Name: value.Name,
			Code:  value.Code,
		}
		users = append(users, user)
	}
	
	return users

}

