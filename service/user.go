package service

import (
	"errors"

	"github.com/pwcong/img-hosting/model"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	Base *BaseService
}

func (ctx *UserService) Register(username string, password string) (model.User, error) {
	db := ctx.Base.DB

	notFound := db.Where("username = ?", username).First(&model.User{}).RecordNotFound()
	if !notFound {
		return model.User{}, errors.New("user existed")
	}

	salt, err := uuid.NewV1()
	if err != nil {
		return model.User{}, err
	}
	password = uuid.NewV5(salt, password).String()

	var user model.User = model.User{
		Username:     username,
		Password:     password,
		PasswordSalt: salt.String(),
	}

	db.Create(&user)

	return user, nil
}

func (ctx *UserService) Login(username string, password string) (model.User, error) {
	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("username = ?", username).First(&user).RecordNotFound()
	if notFound {
		return model.User{}, errors.New("user not existed")
	}

	salt, err := uuid.FromString(user.PasswordSalt)
	if err != nil {
		return model.User{}, err
	}
	_password := uuid.NewV5(salt, password).String()

	if _password != user.Password {
		return model.User{}, errors.New("incorrect password")
	}

	return user, nil
}
