package router

import (
	"github.com/aungkoko1234/tickermaster_backend/controller"
	"github.com/aungkoko1234/tickermaster_backend/middleware"
	repository "github.com/aungkoko1234/tickermaster_backend/repository/user"
	service "github.com/aungkoko1234/tickermaster_backend/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitUserRouters(db *gorm.DB, route *gin.RouterGroup,validate * validator.Validate ) {
		 	//init repository
			 usersRepository := repository.NewUsersRepositoryImpl(db)

			 //init service
			 usersService := service.NewUsersServiceImpl(usersRepository, validate)

			 userController:= controller.NewUserController(usersService)

			 userRouter := route.Group("/users")
			 userRouter.Use(middleware.JwtAuthMiddleware())

		 
		 
			 userRouter.GET("",userController.FindAll)
			 userRouter.GET("/:userId", userController.FindById)
			 userRouter.POST("", userController.Create)
			 userRouter.PATCH("/:userId", userController.Update)
			 userRouter.DELETE("/:userId", userController.Delete)
		
}