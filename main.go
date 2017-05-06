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
	if Init.Config.Server.Middlewares.Gzip.Active {

		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: Init.Config.Server.Middlewares.Gzip.Level,
		}))
	}

	// initialize log middleware configuration
	if Init.Config.Server.Middlewares.Log.Active {

		if Init.Config.Server.Middlewares.Log.Output == "file" {

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
				Format: Init.Config.Server.Middlewares.Log.Format + "\n",
			}))

		} else if Init.Config.Server.Middlewares.Log.Output == "stdout" {
			e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: Init.Config.Server.Middlewares.Log.Format + "\n",
			}))
		}
	}

	// initialize cors middleware configuration
	if Init.Config.Server.Middlewares.Cors.Active {

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: Init.Config.Server.Middlewares.Cors.AllowOrigins,
		}))
	}

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
