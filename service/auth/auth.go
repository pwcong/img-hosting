package auth

import (
	"fmt"
	"time"

	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	Init "github.com/pwcong/img-hosting/init"
)

type UserClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func GenerateTokenStringWithUserClaims(uid string) (error, string) {

	userClaims := UserClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + Init.Config.Server.Jwt.ExpiredTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(Init.Config.Server.Jwt.SigningKey))

	return err, tokenString
}

func CheckTokenStringWithUserClaims(tokenString string) (error, string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Init.Config.Server.Jwt.SigningKey), nil
	})

	if err != nil {
		return err, ""
	}

	if userClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		uid, ok := userClaims["uid"].(string)
		if ok {
			return nil, uid
		}

	}

	return errors.New(""), ""
}
