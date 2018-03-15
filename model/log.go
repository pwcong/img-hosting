package model

import "time"

type Log struct {
	ID        int       `gorm:"type:int;auto_increment;primary_key;not null"`
	IP        string    `gorm:"type:varchar(255);not null"`
	Action    string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}
