package main

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/aungkoko1234/tickermaster_backend/config"
	"github.com/aungkoko1234/tickermaster_backend/controller"
	"github.com/aungkoko1234/tickermaster_backend/model"
	repository "github.com/aungkoko1234/tickermaster_backend/repository/user"
	"github.com/aungkoko1234/tickermaster_backend/router"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
)


func main() {
	//db
    db := config.DatabaseConnection()

	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	//init repository
	usersRepository := repository.NewUsersRepositoryImpl(db)

	//init service
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	//init controller 
	usersController := controller.NewUserController(usersService)
	authController := controller.NewAuthController(usersService)

	//init router

	routes := router.NewRouter(usersController,authController)



	server := &http.Server{
		Addr:  ":8888",
		Handler: routes,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
   
}