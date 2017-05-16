package img

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	Init "github.com/pwcong/img-hosting/init"
	AuthService "github.com/pwcong/img-hosting/service/auth"
	ImgService "github.com/pwcong/img-hosting/service/img"
	stringutils "github.com/pwcong/img-hosting/utils/string"
)

const (
	URL_API_IMG_UPLOAD        = "/img/upload"
	URL_API_IMG_PUBLIC_PREFIX = "/public"
	URL_API_IMG_PUBLIC        = URL_API_IMG_PUBLIC_PREFIX + "/:year/:month/:day/:storename"
	URL_API_IMG_LIST          = "/img/list"
)

type ImgListItemJSON struct {
	Filename  string `json:"filename"`
	Year      string `json:"year"`
	Month     string `json:"month"`
	Day       string `json:"day"`
	Storename string `json:"storename"`
	Path      string `json:"path"`
	URL       string `json:"url"`
}

type ImgListJSONResponse struct {
	Images []ImgListItemJSON `json:"imgs"`
	JSONResponse
}

type ImgJSONResponse struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	URL      string `json:"url"`
	JSONResponse
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

var PublicIMGURLPrefix string

func Upload(c echo.Context) error {

	uid := c.FormValue("uid")

	img, err := c.FormFile("img")

	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			err.Error(),
		})
	}
	contentType := img.Header.Get("Content-Type")

	for _, ext := range Init.Config.Server.Image.SupportExts {

		if ext == contentType {

			src, err := img.Open()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{
					http.StatusInternalServerError,
					err.Error(),
				})
			}

			defer src.Close()

			data, err := ioutil.ReadAll(src)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{http.StatusInternalServerError, err.Error()})
			}
			filename := stringutils.PickFilenameFromContentDisposition(img.Header.Get("Content-Disposition"))

			err, path := ImgService.SaveImage(uid, filename, contentType, data)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, JSONResponse{
					http.StatusInternalServerError,
					err.Error(),
				})
			}

			return c.JSON(http.StatusOK, ImgJSONResponse{
				Filename: filename,
				Path:     path,
				URL:      PublicIMGURLPrefix + path,
				JSONResponse: JSONResponse{
					Code:    http.StatusOK,
					Message: "",
				},
			})
		}
	}

	return c.JSON(http.StatusUnsupportedMediaType, JSONResponse{
		http.StatusUnsupportedMediaType,
		"",
	})
}

func Get(c echo.Context) error {

	err, contentType, data := ImgService.LoadImage(c.Param("year"), c.Param("month"), c.Param("day"), c.Param("storename"))

	if err != nil {
		c.JSON(http.StatusNotFound, JSONResponse{
			http.StatusNotFound,
			err.Error(),
		})
	}

	return c.Blob(http.StatusOK, contentType, data)
}

func List(c echo.Context) error {

	token := c.FormValue("token")

	err, uid := AuthService.CheckTokenStringWithUserClaims(token)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, JSONResponse{
			http.StatusUnauthorized,
			err.Error(),
		})
	}

	if uid == "" {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			"uid can not be null",
		})
	}

	err, imgs := ImgService.GetImagesByUID(uid)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, JSONResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
	}

	var imgListItems []ImgListItemJSON

	for _, img := range imgs {
		imgListItems = append(imgListItems, ImgListItemJSON{
			Filename:  img.Filename,
			Year:      img.Year,
			Month:     img.Month,
			Day:       img.Day,
			Storename: img.Storename,
			Path:      img.Path,
			URL:       PublicIMGURLPrefix + img.Path,
		})
	}

	return c.JSON(http.StatusOK, ImgListJSONResponse{
		Images: imgListItems,
		JSONResponse: JSONResponse{
			Code:    http.StatusOK,
			Message: "",
		},
	})
}

func init() {

	PublicIMGURLPrefix = "http://" + Init.Config.Server.Domain + URL_API_IMG_PUBLIC_PREFIX

}
