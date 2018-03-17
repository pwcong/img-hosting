package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/controller"
	"github.com/pwcong/img-hosting/middleware"
	"github.com/pwcong/img-hosting/service"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	e.Static("/", "view/dist")
	e.Static("/public", "public")
	e.Static("/doc", "doc")

	authMiddleware := middleware.AuthMiddleware{Conf: conf}

	baseService := &service.BaseService{Conf: conf, DB: db}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	indexController := &controller.IndexController{Base: baseController}
	imgController := &controller.ImgController{Base: baseController}
	userController := &controller.UserController{Base: baseController}

	e.GET("/", indexController.Default)
	e.POST("/img/upload", imgController.Upload, authMiddleware.OptionalAuthToken)
	e.POST("/img/remove", imgController.RemovePrivateImages, authMiddleware.AuthToken)
	e.GET("/imgs", imgController.GetPrivateImages, authMiddleware.AuthToken)

	e.POST("/user/login", userController.Login)
	e.POST("/user/register", userController.Register)
	e.POST("/user/check", userController.CheckToken, authMiddleware.AuthToken)

}
