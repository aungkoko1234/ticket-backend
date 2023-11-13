package model

import "gorm.io/gorm"

type Countries struct {
	gorm.Model
	ID       string `gorm:"type:varchar(100);primary_key"`
	Name     string `gorm:"type:varchar(255);not null;"`
	Code     string `gorm:"type:varchar(3); not null;unique"`
}