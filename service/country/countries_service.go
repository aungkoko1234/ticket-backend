package country_service

import (
	"github.com/aungkoko1234/tickermaster_backend/data/request"
	"github.com/aungkoko1234/tickermaster_backend/data/response"
)

type CountriesService interface {
	Create(country request.CreateCountryRequest)
	Update(country request.UpdateCountryRequest)
	Delete (countryId string)
	FindById(countryId string) response.CountryResponse
	FindAll() []response.CountryResponse
}