package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/controller"
	"github.com/pwcong/img-hosting/service"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	e.Static("/", "view/dist")
	e.Static("/public", "public")

	baseService := &service.BaseService{Conf: conf, DB: db}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	indexController := &controller.IndexController{Base: baseController}
	imgController := &controller.ImgController{Base: baseController}

	e.GET("/", indexController.Default)
	e.POST("/img/upload", imgController.Upload)

}
