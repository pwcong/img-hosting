package router

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	e.Static("/", "view")
	// e.GET("/", controller.Index.Default)

}
