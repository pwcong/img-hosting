package model

type User struct {
	Uid string `gorm:"type:varchar(255);primary_key"`
	Pwd string `gorm:"type:varchar(255);not null"`
}
