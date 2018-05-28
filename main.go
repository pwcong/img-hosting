package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/controller"
	"github.com/pwcong/img-hosting/db"
	"github.com/pwcong/img-hosting/middleware"
	"github.com/pwcong/img-hosting/model"
	"github.com/pwcong/img-hosting/router"
	"github.com/pwcong/img-hosting/utils"
)

func initMiddlewares(e *echo.Echo, conf *config.Config) {
	middleware.Init(e, conf)
}

func initRoutes(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	router.Init(e, conf, db)

}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&model.Img{}, &model.User{}, &model.Log{})
}

func main() {

	// 初始化配置
	conf, err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	orm := db.ORM{Name: "mysql"}

	dbConfig, ok := conf.Databases["mysql"]
	if !ok {
		log.Fatal("Can not load configuration of MySQL")
	}

	connectionURL := dbConfig.Username + ":" + dbConfig.Password + "@" + "tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" + dbConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"

	orm.Open(connectionURL)

	defer orm.Close()

	initDB(orm.DB)

	e := echo.New()

	// 全局错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusOK, controller.BaseResponseJSON{
			Success: false,
			Code:    controller.STATUS_ERROR,
			Message: err.Error(),
		})
	}

	// 初始化中间件
	initMiddlewares(e, &conf)
	// 初始化路由
	initRoutes(e, &conf, orm.DB)

	// 运行服务
	if conf.Server.Port == 80 {
		e.Logger.Fatal(e.Start(conf.Server.Host))
	} else {
		e.Logger.Fatal(e.Start(conf.Server.Host + ":" + strconv.Itoa(conf.Server.Port)))
	}

}

func init() {
	// 初始化目录
	err := utils.MkdirIFNotExist("public")
	if err != nil {
		log.Fatal(err)
	}

}
