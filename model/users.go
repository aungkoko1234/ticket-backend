package model

import (
	"fmt"
	"html"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID string `gorm:"type:varchar(100);primary_key"`
	Name  string `gorm:"type:varchar(255);not null;"`
	Email string `gorm:"type:varchar(255); not null;unique"`
	Password string `gorm:"type:varchar(255); not null;"`
}

func (user *Users) BeforeCreate(db *gorm.DB) error {
	uuid, err := gonanoid.New()
	fmt.Println("uuid",uuid)
	if err != nil {
	 return err
	}

	user.ID = uuid

	return nil
   }
   

func (u *Users) BeforeSave(*gorm.DB) error{
    // id, err :=  gonanoid.New()

	// fmt.Println("nanoid",id)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}