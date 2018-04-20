package middleware

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pwcong/img-hosting/config"
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
			return errors.New("lack of token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(ctx.Conf.Auth.Secret), nil
		})

		if err != nil {
			return errors.New("invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id, ok := claims["id"]
			if !ok {
				return errors.New("invalid token")
			} else {
				c.Set("id", id)
				c.Set("token", tokenString)
			}
		} else {
			return errors.New("invalid token")
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
