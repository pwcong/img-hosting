package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pwcong/img-hosting/config"
)

func Init(e *echo.Echo, conf *config.Config) {
	middlewaresConfig := conf.Middlewares

	loggerConfig, ok := middlewaresConfig["logger"]
	if ok && loggerConfig.Active {
		e.Use(middleware.Logger())
	}

	corsConfig, ok := middlewaresConfig["cors"]
	if ok && corsConfig.Active {
		e.Use(middleware.CORS())

	}
}
