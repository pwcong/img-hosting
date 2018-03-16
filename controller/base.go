package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/service"
)

const (
	STATUS_OK    = 20000
	STATUS_ERROR = 40000
)

type BaseResponseJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func BaseResponse(c echo.Context, code int, message string, payload interface{}) error {

	return c.JSON(http.StatusOK, BaseResponseJSON{
		Code:    code,
		Message: message,
		Payload: payload,
	})

}

type BaseController struct {
	Conf    *config.Config
	Service *service.BaseService
}
