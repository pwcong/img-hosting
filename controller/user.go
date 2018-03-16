package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/service"
)

type UserForm struct {
	Username string `json:"username" xml:"username" form:"username" query:"username"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

type UserController struct {
	Base *BaseController
}

func (ctx *UserController) Login(c echo.Context) error {

	service := service.UserService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "login")

	uf := new(UserForm)
	if err := c.Bind(uf); err != nil {
		return BaseResponse(c, STATUS_ERROR, "invalid params", struct{}{})
	}

	if uf.Username == "" || uf.Password == "" {
		return BaseResponse(c, STATUS_ERROR, "params not enough", struct{}{})
	}

	user, err := service.Login(uf.Username, uf.Password)
	if err != nil {
		return BaseResponse(c, STATUS_ERROR, err.Error(), struct{}{})
	}

	now := time.Now()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	expiredTime := ctx.Base.Conf.Auth.ExpiredTime
	claims["exp"] = now.Add(time.Second * time.Duration(expiredTime)).Unix()

	t, err := token.SignedString([]byte(ctx.Base.Conf.Auth.Secret))
	if err != nil {
		BaseResponse(c, STATUS_ERROR, err.Error(), struct{}{})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = now.Add(time.Duration(expiredTime/3600) * time.Hour)

	c.SetCookie(cookie)

	return BaseResponse(c, STATUS_OK, "login successfully", struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Token:    t,
	})

}

func (ctx *UserController) Register(c echo.Context) error {

	service := service.UserService{Base: ctx.Base.Service}

	service.Base.Info(c.RealIP(), "register")

	uf := new(UserForm)
	if err := c.Bind(uf); err != nil {
		return BaseResponse(c, STATUS_ERROR, "invalid params", struct{}{})
	}

	if uf.Username == "" || uf.Password == "" {
		return BaseResponse(c, STATUS_ERROR, "params not enough", struct{}{})
	}

	user, err := service.Register(uf.Username, uf.Password)
	if err != nil {
		return BaseResponse(c, STATUS_ERROR, err.Error(), struct{}{})
	}

	now := time.Now()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	expiredTime := ctx.Base.Conf.Auth.ExpiredTime
	claims["exp"] = now.Add(time.Second * time.Duration(expiredTime)).Unix()

	t, err := token.SignedString([]byte(ctx.Base.Conf.Auth.Secret))
	if err != nil {
		BaseResponse(c, STATUS_ERROR, err.Error(), struct{}{})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = now.Add(time.Duration(expiredTime/3600) * time.Hour)

	c.SetCookie(cookie)

	return BaseResponse(c, STATUS_OK, "register successfully", struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Token:    t,
	})

}
