package controller

import (
	"github.com/labstack/echo/v4"
)

type IndexController struct {
	Base *BaseController
}

func (ctx *IndexController) Default(c echo.Context) error {
	return c.File("view/dist/index.html")
}
