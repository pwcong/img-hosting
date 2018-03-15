package service

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/model"
)

type BaseService struct {
	Conf *config.Config
	DB   *gorm.DB
}

func (ctx *BaseService) Log(ip string, action string) error {

	now := time.Now()

	ctx.DB.Create(model.Log{
		IP:        ip,
		Action:    action,
		CreatedAt: now,
		UpdatedAt: now,
	})

	return nil

}
