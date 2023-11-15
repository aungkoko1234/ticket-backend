package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Countries struct {
	gorm.Model
	ID       string `gorm:"type:varchar(100);primary_key"`
	Name     string `gorm:"type:varchar(255);not null;"`
	Code     string `gorm:"type:varchar(3); not null;unique"`
}

func (country *Countries) BeforeCreate(db *gorm.DB) error {
	uuid, err := gonanoid.New()
	if err != nil {
	 return err
	}

	country.ID = uuid

	return nil
   }