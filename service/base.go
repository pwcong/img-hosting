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

func (ctx *BaseService) Log(ip string, action string) error {

	ctx.DB.Create(model.Log{
		IP:     ip,
		Action: action,
	})

	return nil

}
