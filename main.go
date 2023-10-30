package main

import (
	"net/http"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/aungkoko1234/tickermaster_backend/config"
	"github.com/aungkoko1234/tickermaster_backend/model"
	"github.com/aungkoko1234/tickermaster_backend/router"
)


func main() {
	//db
    db := config.DatabaseConnection()

	db.Table("users").AutoMigrate(&model.Users{})

	//init router

	routes := setUpRouters()



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

func setUpRouters() *gin.Engine{
	db := config.DatabaseConnection()

	validate := validator.New()

	service := gin.Default()

	service.GET("",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome Home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	route := service.Group("/api")

	route.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	route.Use(helmet.Default())
	route.Use(gzip.Gzip(gzip.BestCompression))

	db.Table("users").AutoMigrate(&model.Users{})

	router.InitUserRouters(db,route,validate)
	router.IntiAuthRouters(db,route,validate)

	return service

}