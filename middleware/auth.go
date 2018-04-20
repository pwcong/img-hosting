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

		cookie, err := c.Request().Cookie("Token")

		var tokenString string
		if err == nil {
			fmt.Println(cookie.Value)
			tokenString = cookie.Value
		} else {
			tokenString = c.Request().Header.Get("Token")
		}

		if tokenString == "" {
			return BaseResponse(c, false, STATUS_ERROR, "lack of token", struct{}{})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return BaseResponse(c, false, STATUS_ERROR, "invalid token", struct{}{})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, ok := claims["id"]
			if !ok {
				return BaseResponse(c, false, STATUS_ERROR, "invalid token", struct{}{})
			} else {
				c.Set("id", id)
				c.Set("token", tokenString)
			}
		} else {
			return BaseResponse(c, false, STATUS_ERROR, "invalid token", struct{}{})
		}

		return next(c)
	}
}

func (ctx *AuthMiddleware) OptionalAuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString := c.Request().Header.Get("Token")

		if tokenString == "" {
			return next(c)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return next(c)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, ok := claims["id"]
			if !ok {
				return next(c)
			} else {
				c.Set("id", id)
				c.Set("token", tokenString)
			}
		} else {
			return next(c)
		}

		return next(c)
	}
}
