package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/model"
)

type BaseService struct {
	Conf *config.Config
	DB   *gorm.DB
}

const (
	INFO = iota
	WARN
	ERROR
)

func (ctx *BaseService) Log(level uint, ip string, action string) {
	db := ctx.DB

	db.Create(&model.Log{
		Level:  level,
		IP:     ip,
		Action: action,
	})

}

func (ctx *BaseService) Info(ip string, action string) {
	ctx.Log(INFO, ip, action)
}
func (ctx *BaseService) Warn(ip string, action string) {
	ctx.Log(INFO, ip, action)
}
func (ctx *BaseService) Error(ip string, action string) {
	ctx.Log(INFO, ip, action)
}
