package main

import (
	"log"
	"os"
	"strconv"

	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	Init "github.com/pwcong/img-hosting/init"
	"github.com/pwcong/img-hosting/router"
)

func initMiddlewares(e *echo.Echo) {

	// initialize gzip middleware configuration
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: Init.Config.Server.CompressLevel,
	}))

	// initialize log middleware configuration
	if Init.Config.Mode == "prod" {

		logDir := filepath.Join(filepath.Dir(os.Args[0]), "log")
		if _, err := os.Stat(logDir); err != nil {
			err := os.MkdirAll(logDir, 0666)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		logPath := filepath.Join(logDir, "server.log")
		logOutput, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err.Error())
		}

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logOutput,
		}))

	} else {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	}

	// initialize cors middleware configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
}

func initRoutes(e *echo.Echo) {
	router.Init(e)
}

func main() {

	e := echo.New()

	initMiddlewares(e)
	initRoutes(e)

	e.Logger.Fatal(e.Start(Init.Config.Server.Host + ":" + strconv.Itoa(Init.Config.Server.Port)))
}
