package index

import (
	"github.com/labstack/echo"
)

const (
	URL_INDEX = "/"
)

func Default(c echo.Context) error {
	return c.File("view/index.html")
}
