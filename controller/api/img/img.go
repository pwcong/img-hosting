package img

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	Init "github.com/pwcong/img-hosting/init"
	ImgService "github.com/pwcong/img-hosting/service/img"
	stringutils "github.com/pwcong/img-hosting/utils/string"
)

const (
	URL_API_IMG_UPLOAD        = "/img/upload"
	URL_API_IMG_PUBLIC_PREFIX = "/public"
	URL_API_IMG_PUBLIC        = URL_API_IMG_PUBLIC_PREFIX + "/:year/:month/:day/:storename"
)

type ImgJSONResponse struct {
	Code     int    `json:"code"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
	URL      string `json:"url"`
	JSONResponse
}

type JSONResponse struct {
	Message string `json:"message"`
}

var PublicIMGURLPrefix string

func Upload(c echo.Context) error {

	fmt.Println(c.Path())

	uid := c.FormValue("uid")

	img, err := c.FormFile("img")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSONResponse{err.Error()})
	}

	contentType := img.Header.Get("Content-Type")

	for _, ext := range Init.Config.Server.Image.SupportExts {

		if ext == contentType {

			src, err := img.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{err.Error()})
			}

			defer src.Close()

			data, err := ioutil.ReadAll(src)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{err.Error()})
			}
			filename := stringutils.PickFilenameFromContentDisposition(img.Header.Get("Content-Disposition"))

			err, path := ImgService.SaveImage(uid, filename, contentType, data)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{err.Error()})
			}

			return c.JSON(http.StatusOK, ImgJSONResponse{
				Code:     http.StatusOK,
				Filename: filename,
				Path:     path,
				URL:      PublicIMGURLPrefix + path,
			})
		}
	}

	return c.JSON(http.StatusUnsupportedMediaType, JSONResponse{""})
}

func Get(c echo.Context) error {

	err, contentType, data := ImgService.LoadImage(c.Param("year"), c.Param("month"), c.Param("day"), c.Param("storename"))

	if err != nil {
		c.JSON(http.StatusNotFound, JSONResponse{err.Error()})
	}

	return c.Blob(http.StatusOK, contentType, data)
}

func init() {

	PublicIMGURLPrefix = "http://" + Init.Config.Server.Domain + URL_API_IMG_PUBLIC_PREFIX

}
