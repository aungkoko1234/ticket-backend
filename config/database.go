package config

import (
	"fmt"
	"os"

	"github.com/aungkoko1234/tickermaster_backend/helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func DatabaseConnection() *gorm.DB {
	err := godotenv.Load(".env")

	helper.ErrorPanic(err)

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	helper.ErrorPanic(err)

	return db


}