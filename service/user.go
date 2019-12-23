package service

import (
	"errors"

	"img-hosting/model"

	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	Base *BaseService
}

func (ctx *UserService) Register(username string, password string) (model.User, error) {
	db := ctx.Base.DB

	notFound := db.Where("username = ?", username).First(&model.User{}).RecordNotFound()
	if !notFound {
		return model.User{}, errors.New("user is existed")
	}

	salt := uuid.NewV1()

	password = uuid.NewV5(salt, password).String()

	var user model.User = model.User{
		Username:     username,
		Password:     password,
		PasswordSalt: salt.String(),
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		return model.User{}, dbc.Error
	}

	return user, nil
}

func (ctx *UserService) Login(username string, password string) (model.User, error) {
	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("username = ?", username).First(&user).RecordNotFound()
	if notFound {
		return model.User{}, errors.New("user is not existed")
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

func (ctx *UserService) FindByID(id uint) (model.User, error) {
	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("id = ?", id).First(&user).RecordNotFound()
	if notFound {
		return model.User{}, errors.New("user is not existed")
	}

	return user, nil
}
