package main

import (
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
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

	mySQLConfig, ok := conf.Databases["mysql"]
	if !ok {
		log.Fatal("Can not load configuration of MySQL")
	}

	orm := db.ORM{DB: nil, Name: "mysql"}

	orm.Open(
		mySQLConfig.Username,
		mySQLConfig.Password,
		mySQLConfig.Host+":"+strconv.Itoa(mySQLConfig.Port),
		mySQLConfig.DBName)

	defer orm.Close()

	initDB(orm.DB)

	e := echo.New()
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
