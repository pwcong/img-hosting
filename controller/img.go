package controller

import (
	"time"

	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/service"
)

type ImgItem struct {
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ImgController struct {
	Base *BaseController
}

func (ctx *ImgController) Upload(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload image")

	file, err := c.FormFile("img")
	if file == nil || err != nil {
		return BaseResponse(c, STATUS_OK, err.Error(), struct{}{})
	}

	img, err := service.SaveImage(file)
	if err != nil {
		return BaseResponse(c, STATUS_OK, err.Error(), struct{}{})
	}

	return BaseResponse(c, STATUS_OK, "upload successfully", struct {
		URL string `json:"url"`
	}{
		URL: "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
	})

}

func (ctx *ImgController) PrivateUpload(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload private image")

	file, err := c.FormFile("img")
	if file == nil || err != nil {
		return BaseResponse(c, STATUS_OK, err.Error(), struct{}{})
	}

	img, err := service.SavePrivateImage(file, int(c.Get("id").(float64)))
	if err != nil {
		return BaseResponse(c, STATUS_OK, err.Error(), struct{}{})
	}

	return BaseResponse(c, STATUS_OK, "upload successfully", struct {
		URL string `json:"url"`
	}{
		URL: "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
	})

}

func (ctx *ImgController) GetPrivateImages(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "get private images")

	imgs, err := service.GetPrivateImages(int(c.Get("id").(float64)))
	if err != nil {
		return BaseResponse(c, STATUS_OK, err.Error(), struct{}{})
	}

	var _imgs []interface{}

	for _, img := range imgs {
		_imgs = append(_imgs, ImgItem{
			URL:       "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
			CreatedAt: img.CreatedAt,
		})
	}

	res := Pager(c, _imgs)

	return BaseResponse(c, STATUS_OK, "query successfully", res)

}
