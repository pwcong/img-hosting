package controller

import (
	"time"

	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/model"
	"github.com/pwcong/img-hosting/service"
)

type ImgItem struct {
	ID        uint      `json:"id"`
	URL       string    `json:"url"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
}

type ImgForm struct {
	IDs []uint `json:"ids" form:"ids"`
}

type ImgController struct {
	Base *BaseController
}

func (ctx *ImgController) Upload(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload image")

	file, err := c.FormFile("img")
	if file == nil || err != nil {
		return BaseResponse(c, false, STATUS_ERROR, err.Error(), struct{}{})
	}

	var img model.Img

	id := c.Get("id")
	if id == nil {
		img, err = service.SaveImage(file)

	} else {
		img, err = service.SavePrivateImage(file, uint(id.(float64)))
	}

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		ID  uint   `json:"id"`
		URL string `json:"url"`
	}{
		ID:  img.ID,
		URL: "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
	})

}

func (ctx *ImgController) PrivateUpload(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "upload private image")

	file, err := c.FormFile("img")
	if file == nil || err != nil {
		return err
	}

	img, err := service.SavePrivateImage(file, uint(c.Get("id").(float64)))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		ID  uint   `json:"id"`
		URL string `json:"url"`
	}{
		ID:  img.ID,
		URL: "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
	})

}

func (ctx *ImgController) GetPrivateImages(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "get private images")

	imgs, err := service.GetPrivateImages(uint(c.Get("id").(float64)))
	if err != nil {
		return err
	}

	var _imgs []interface{}

	for _, img := range imgs {
		_imgs = append(_imgs, ImgItem{
			ID:        img.ID,
			Filename:  img.Filename,
			URL:       "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.Symbol + "." + img.ExtName,
			CreatedAt: img.CreatedAt,
		})
	}

	res := Pager(c, _imgs)

	return BaseResponse(c, true, STATUS_OK, "query successfully", res)

}

func (ctx *ImgController) RemovePrivateImages(c echo.Context) error {

	service := service.ImgService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "delete private images")

	imgForm := new(ImgForm)

	if err := c.Bind(imgForm); err != nil {
		return err
	}

	_, err := service.RemovePrivateImages(uint(c.Get("id").(float64)), imgForm.IDs)
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "delete successfully", struct{}{})

}
