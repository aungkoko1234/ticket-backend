package model

import (
	"fmt"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id int `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(255);not null;"`
	Email string `gorm:"type:varchar(255); not null;unique"`
	Password string `gorm:"type:varchar(255); not null;"`
}

func (u *Users) BeforeSave(*gorm.DB) error{
	fmt.Println("password plain",u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println("hashing",string(hashedPassword))
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}