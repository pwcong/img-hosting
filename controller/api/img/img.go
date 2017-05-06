package img

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"strconv"

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

type JSONResponse struct {
	Code     int
	Filename string
	Path     string
	URL      string
}

var PublicIMGURLPrefix string

func Upload(c echo.Context) error {

	fmt.Println(c.Path())

	uid := c.FormValue("uid")

	img, err := c.FormFile("img")

	if err != nil {
		return err
	}

	contentType := img.Header.Get("Content-Type")

	for _, ext := range Init.Config.Server.Image.SupportExts {

		if ext == contentType {

			src, err := img.Open()
			if err != nil {
				return err
			}

			defer src.Close()

			data, err := ioutil.ReadAll(src)
			if err != nil {
				return err
			}
			filename := stringutils.PickFilenameFromContentDisposition(img.Header.Get("Content-Disposition"))

			err, path := ImgService.SaveImage(uid, filename, data)

			if err != nil {
				return c.String(http.StatusInternalServerError, "")
			}

			return c.JSON(http.StatusOK, JSONResponse{
				Code:     http.StatusOK,
				Filename: filename,
				Path:     path,
				URL:      PublicIMGURLPrefix + path,
			})
		}
	}

	return c.String(http.StatusUnsupportedMediaType, "")
}

func Get(c echo.Context) error {

	err, data := ImgService.LoadImage(c.Param("year"), c.Param("month"), c.Param("day"), c.Param("storename"))

	if err != nil {
		c.String(http.StatusNotFound, "")
	}

	return c.Blob(http.StatusOK, http.DetectContentType(data), data)
}

func init() {

	if Init.Config.Server.Port == 80 {
		PublicIMGURLPrefix = "http://" + Init.Config.Server.Domain + URL_API_IMG_PUBLIC_PREFIX
	} else {
		PublicIMGURLPrefix = "http://" +
			Init.Config.Server.Domain +
			":" +
			strconv.Itoa(Init.Config.Server.Port) +
			URL_API_IMG_PUBLIC_PREFIX
	}

}
