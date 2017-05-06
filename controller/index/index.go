package index

import (
	"net/http"

	"github.com/labstack/echo"
)

func Default(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("year")+c.Param("month")+c.Param("day")+c.Param("img"))
}
