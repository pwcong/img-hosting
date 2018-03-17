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

type BaseController struct {
	Conf    *config.Config
	Service *service.BaseService
}

type BaseResponseJSON struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func BaseResponse(c echo.Context, success bool, code int, message string, payload interface{}) error {

	return c.JSON(http.StatusOK, BaseResponseJSON{
		Success: success,
		Code:    code,
		Message: message,
		Payload: payload,
	})

}

type Page struct {
	Data      []interface{} `json:"data"`
	PageSize  int           `json:"page_size"`
	PageNo    int           `json:"page_no"`
	TotalSize int           `json:"total_size"`
	TotalNo   int           `json:"total_no"`
}

type PageForm struct {
	PageSize int `json:"page_size" xml:"page_size" form:"page_size" query:"page_size"`
	PageNo   int `json:"page_no" xml:"page_no" form:"page_no" query:"page_no"`
}

func Pager(c echo.Context, data []interface{}) interface{} {

	l := len(data)

	var pageSize int
	var pageNo int

	pf := new(PageForm)

	if err := c.Bind(pf); err != nil {
		return Page{
			Data:      data,
			PageSize:  l,
			PageNo:    1,
			TotalSize: 1,
			TotalNo:   1,
		}
	}

	pageSize = pf.PageSize
	pageNo = pf.PageNo

	if pageSize <= 0 && pageNo <= 0 {

		pageSize = l
		pageNo = 1

	} else {

		if pageSize <= 0 {
			pageSize = 20
		}

		if pageNo <= 0 {
			pageNo = 1
		}
	}

	start := (pageNo - 1) * pageSize
	end := pageNo * pageSize

	if start >= l {
		start = l
	}

	if end >= l {
		end = l
	}

	var totalNo int
	if l%pageSize == 0 {
		totalNo = l / pageSize
	} else {
		totalNo = l/pageSize + 1
	}

	return Page{
		Data:      data[start:end],
		PageSize:  pageSize,
		PageNo:    pageNo,
		TotalSize: l,
		TotalNo:   totalNo,
	}

}
