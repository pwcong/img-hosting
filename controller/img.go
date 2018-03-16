package controller

import (
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/service"
)

type ImgController struct {
	Base *BaseController
}

func (ctx *ImgController) Upload(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	file, err := c.FormFile("img")
	if err != nil {
		return err
	}

	url, err := service.SaveImage(file)
	if err != nil {
		return err
	}

	return BaseResponse(c, STATUS_OK, "upload successfully", struct {
		URL string `json:"url"`
	}{
		URL: url,
	})

}
