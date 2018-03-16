package model

import (
	"github.com/jinzhu/gorm"
)

type Img struct {
	gorm.Model
	Symbol   string  `gorm:"type:varchar(255);not null;unique"`
	Filename string  `gorm:"type:varchar(255);not null"`
	Year     string  `gorm:"type:varchar(255);not null"`
	Month    string  `gorm:"type:varchar(255);not null"`
	Date     string  `gorm:"type:varchar(255);not null"`
	ExtName  string  `gorm:"type:varchar(255);not null"`
	Users    []*User `gorm:"many2many:user_imgs;association_foreignkey:id;foreignkey:id"`
}
