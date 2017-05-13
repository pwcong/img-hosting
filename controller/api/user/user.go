package user

import (
	"net/http"

	"github.com/labstack/echo"

	AuthService "github.com/pwcong/img-hosting/service/auth"
	UserService "github.com/pwcong/img-hosting/service/user"
)

const (
	URL_API_USER_REGISTER = "/user/register"
	URL_API_USER_LOGIN    = "/user/login"
	URL_API_USER_UPDATE   = "/user/update"
)

type UserJSONResponse struct {
	Token string `json:"token"`
	JSONResponse
}

type JSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Register(c echo.Context) error {

	uid := c.FormValue("uid")
	pwd := c.FormValue("pwd")

	if uid == "" || pwd == "" {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			"uid and pwd can not be empty",
		})
	}

	err, uid := UserService.Add(uid, pwd)

	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			err.Error(),
		})

	}

	err, tokenString := AuthService.GenerateTokenStringWithUserClaims(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSONResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
	}

	return c.JSON(http.StatusOK, UserJSONResponse{
		Token: tokenString,
		JSONResponse: JSONResponse{
			Code:    http.StatusOK,
			Message: "",
		},
	})
}

func Login(c echo.Context) error {

	uid := c.FormValue("uid")
	pwd := c.FormValue("pwd")

	if uid == "" || pwd == "" {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			"uid and pwd can not be empty",
		})
	}

	err, user := UserService.Find(uid, pwd)

	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			err.Error(),
		})
	}

	err, tokenString := AuthService.GenerateTokenStringWithUserClaims(user.Uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSONResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
	}

	return c.JSON(http.StatusOK, UserJSONResponse{
		Token: tokenString,
		JSONResponse: JSONResponse{
			Code:    http.StatusOK,
			Message: "",
		},
	})
}

func UpdatePassword(c echo.Context) error {

	uid := c.FormValue("uid")
	srcPwd := c.FormValue("srcPwd")
	newPwd := c.FormValue("newPwd")

	if uid == "" || srcPwd == "" || newPwd == "" {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			"uid and srcPwd and newPwd can not be empty",
		})
	}

	err, _ := UserService.Find(uid, srcPwd)

	if err != nil {
		return c.JSON(http.StatusBadRequest, JSONResponse{
			http.StatusBadRequest,
			err.Error(),
		})
	}

	err = UserService.UpdatePassword(uid, srcPwd, newPwd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, JSONResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)

}
