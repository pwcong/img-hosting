package model

type User struct {
	Uid string `gorm:"type:varchar(100);primary_key"`
	Pwd string `gorm:"type:varchar(100);not null"`
}
