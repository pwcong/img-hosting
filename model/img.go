package model

import "time"

type Img struct {
	ID        string    `gorm:"type:varchar(255);primary_key;not null"`
	Filename  string    `gorm:"type:varchar(255);not null"`
	Year      string    `gorm:"type:varchar(255);not null"`
	Month     string    `gorm:"type:varchar(255);not null"`
	Date      string    `gorm:"type:varchar(255);not null"`
	ExtName   string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}
