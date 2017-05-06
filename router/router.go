package router

import (
	"github.com/labstack/echo"
	IndexController "github.com/pwcong/img-hosting/controller/index"
)

func Init(e *echo.Echo) {

	e.Static("/", "view")

	e.GET("/:year/:month/:day/:img", IndexController.Default)

}
