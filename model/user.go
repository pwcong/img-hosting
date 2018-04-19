package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(255);unique_index;not null;"`
	Password     string `gorm:"type:varchar(255);not null;"`
	PasswordSalt string `gorm:"type:varchar(255);not null;"`
	Imgs         []*Img `gorm:"many2many:user_imgs;"`
}
