package controller

import (
	"github.com/pwcong/img-hosting/config"
	"github.com/pwcong/img-hosting/service"
)

type BaseController struct {
	Conf    *config.Config
	Service *service.BaseService
}

type BaseJSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
