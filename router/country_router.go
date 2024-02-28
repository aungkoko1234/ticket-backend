package router

import (
	"github.com/aungkoko1234/tickermaster_backend/controller"
	"github.com/aungkoko1234/tickermaster_backend/middleware"
	country_repository "github.com/aungkoko1234/tickermaster_backend/repository/country"
	country_service "github.com/aungkoko1234/tickermaster_backend/service/country"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitCountryRouters(db *gorm.DB, route *gin.RouterGroup,validate *validator.Validate ) {
		 	//init repository
			 countryRepository := country_repository.NewCountriesRepositoryImpl(db)

			 //init service
			 countriesService := country_service.NewCountriesServiceImpl(countryRepository, validate)

			 countryController:= controller.NewCountryController(countriesService)

			 countryRouter := route.Group("/countries")
			 countryRouter.Use(middleware.JwtAuthMiddleware())

		 
		 
			 countryRouter.GET("",countryController.FindAll)
			 countryRouter.GET("/:countryId", countryController.FindById)
			 countryRouter.POST("", countryController.Create)
			 countryRouter.PATCH("/:countryId", countryController.Update)
			 countryRouter.DELETE("/:countryId", countryController.Delete)
		
}