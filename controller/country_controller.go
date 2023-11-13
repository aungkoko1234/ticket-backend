package controller

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/data/request"
	country_service "github.com/aungkoko1234/tickermaster_backend/service/country"
	"github.com/aungkoko1234/tickermaster_backend/utils"
	"github.com/gin-gonic/gin"
)

type CountryController struct {
	countriesService  country_service.CountriesService
}

func NewCountryController(service country_service.CountriesService) *CountryController{
	return  &CountryController{
		countriesService : service,
	}
}

func (controller *CountryController) Create(ctx *gin.Context) {
	createCountryRequest := request.CreateCountryRequest{}

	if err := ctx.ShouldBindJSON(&createCountryRequest); err != nil {
		message := utils.ParseError(err)
        utils.ValidationErrorReponse(ctx, http.StatusBadRequest, http.MethodPost, message)
		return
	}

	controller.countriesService.Create(createCountryRequest)

	utils.ApiResponse(ctx, "Create Country Success", http.StatusOK, http.MethodPost, gin.H{"message": "Success"})
}

func (controller *CountryController) Update(ctx *gin.Context) {
	updateCountryRequest := request.UpdateCountryRequest{}

	if err := ctx.ShouldBindJSON(&updateCountryRequest); err != nil {
		message := utils.ParseError(err)
		utils.ValidationErrorReponse(ctx, http.StatusBadRequest, http.MethodPut, message)
	}

	controller.countriesService.Update(updateCountryRequest)
	utils.ApiResponse(ctx, "Update Country Success", http.StatusOK, http.MethodPost, gin.H{"message": "Success"})
}

func (controller *CountryController) Delete(ctx *gin.Context) {
	countryId := ctx.Param("countryId")
	controller.countriesService.Delete(countryId)
	utils.ApiResponse(ctx, "Delete Country Success", http.StatusOK, http.MethodPost, gin.H{"message": "Success"})
}

func (controller *CountryController) FindById(ctx *gin.Context) {
	countryId := ctx.Param("countryId")
	countryResponse := controller.countriesService.FindById(countryId)
	utils.ApiResponse(ctx,"Country Detail",http.StatusOK, http.MethodGet,countryResponse)
}

func (controller *CountryController) FindAll(ctx *gin.Context) {
     countryResponse := controller.countriesService.FindAll()
	 utils.ApiResponse(ctx,"Country List",http.StatusOK, http.MethodGet,countryResponse)
}   