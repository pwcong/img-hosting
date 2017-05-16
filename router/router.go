package router

import (
	"github.com/labstack/echo"
	ImgController "github.com/pwcong/img-hosting/controller/api/img"
	UserController "github.com/pwcong/img-hosting/controller/api/user"
	IndexController "github.com/pwcong/img-hosting/controller/index"
)

func Init(e *echo.Echo) {

	e.Static("/", "view")

	e.GET(IndexController.URL_INDEX, IndexController.Default)

	e.POST(ImgController.URL_API_IMG_UPLOAD, ImgController.Upload)
	e.GET(ImgController.URL_API_IMG_PUBLIC, ImgController.Get)
	e.POST(ImgController.URL_API_IMG_LIST, ImgController.List)

	e.POST(UserController.URL_API_USER_LOGIN, UserController.Login)
	e.POST(UserController.URL_API_USER_REGISTER, UserController.Register)
	e.POST(UserController.URL_API_USER_UPDATE, UserController.UpdatePassword)

}
