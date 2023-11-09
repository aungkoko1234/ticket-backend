package router

import (
	"github.com/aungkoko1234/tickermaster_backend/controller"
	repository "github.com/aungkoko1234/tickermaster_backend/repository/user"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func IntiAuthRouters(db *gorm.DB, route *gin.RouterGroup,validate * validator.Validate ) {
	 	//init repository
	usersRepository := repository.NewUsersRepositoryImpl(db)

	//init service
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	authController := controller.NewAuthController(usersService)

	authRouter := route.Group("/auth")
	authRouter.POST("login",authController.Login)
	authRouter.POST("register",authController.Register)
}