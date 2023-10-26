package router

import (
	"net/http"

	"github.com/aungkoko1234/tickermaster_backend/controller"
	"github.com/aungkoko1234/tickermaster_backend/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController,authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	service.GET("",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome Home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("login",authController.Login)
	
	userRouter := router.Group("/users")
	userRouter.Use(middleware.JwtAuthMiddleware())


	userRouter.GET("",userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)


	return service
}