package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
	. "github.com/pwcong/img-hosting/controller"
)

type AuthMiddleware struct {
	Conf *config.Config
}

func (ctx *AuthMiddleware) AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("Token")

		if tokenString == "" {
			return BaseResponse(c, STATUS_ERROR, "lack of token", struct{}{})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return BaseResponse(c, STATUS_ERROR, "invalid token", struct{}{})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, ok := claims["id"]
			if !ok {
				return BaseResponse(c, STATUS_ERROR, "invalid token", struct{}{})
			} else {
				c.Set("id", id)
			}
		} else {
			return BaseResponse(c, STATUS_ERROR, "invalid token", struct{}{})
		}

		return next(c)
	}
}
