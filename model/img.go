package model

import (
	"time"
)

type Img struct {
	ID        uint      `gorm:"AUTO_INCREMENT;primary_key;unsigned"`
	Data      []byte    `gorm:"type:mediumblob;not null"`
	Filename  string    `gorm:"type:varchar(255);not null"`
	Year      string    `gorm:"type:varchar(255);not null"`
	Month     string    `gorm:"type:varchar(255);not null"`
	Day       string    `gorm:"type:varchar(255);not null"`
	Storename string    `gorm:"type:varchar(255);not null"`
	Path      string    `gorm:"type:varchar(255);not null"`
	Uid       string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
}
