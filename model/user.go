package model

import "time"

type User struct {
	ID           int       `gorm:"type:int;auto_increment;primary_key;not null"`
	Username     string    `gorm:"type:varchar(255);unique_index;not null"`
	Password     string    `gorm:"type:varchar(255);not null"`
	PasswordSalt string    `gorm:"type:varchar(255);not null"`
	Imgs         []Img     `gorm:"many2many:user_imgs"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}
