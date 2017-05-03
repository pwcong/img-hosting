package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type IndexController struct{}

var Index IndexController

func (ctl *IndexController) Default(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
