package middleware

import (
	"img-hosting/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
