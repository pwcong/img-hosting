package router

import (
	"github.com/labstack/echo"
	ImgController "github.com/pwcong/img-hosting/controller/api/img"
)

func Init(e *echo.Echo) {

	e.Static("/", "view")

	e.POST(ImgController.URL_API_IMG_UPLOAD, ImgController.Upload)

	e.GET(ImgController.URL_API_IMG_PUBLIC, ImgController.Get)

}
