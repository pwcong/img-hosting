package model

import "github.com/jinzhu/gorm"

type Log struct {
	gorm.Model
	Level  uint   `gorm:"not null"`
	IP     string `gorm:"type:varchar(255);not null"`
	Action string `gorm:"type:text;not null"`
}
