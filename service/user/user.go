package user

import (
	"errors"

	"github.com/pwcong/img-hosting/db"
	"github.com/pwcong/img-hosting/model"
)

func Add(uid string, pwd string) (error, string) {

	user := model.User{
		Uid: uid,
		Pwd: pwd,
	}

	notFound := db.MySQL.DB.First(&model.User{}, "uid = ?", uid).RecordNotFound()

	if !notFound {
		return errors.New("user is existed"), ""
	}

	db.MySQL.DB.Create(&user)

	notFound = db.MySQL.DB.First(&model.User{}, "uid = ? AND pwd = ?", uid, pwd).RecordNotFound()

	if !notFound {
		return nil, uid
	}

	return errors.New("unknown error"), ""
}

func Find(uid string, pwd string) (error, model.User) {

	var user model.User

	notFound := db.MySQL.DB.First(&user, "uid = ? AND pwd = ?", uid, pwd).RecordNotFound()

	if notFound {
		return errors.New("invalid uid or pwd"), model.User{}
	}

	return nil, user
}

func Delete(uid string, pwd string) (error, string) {

	err, user := Find(uid, pwd)

	if err == nil {

		db.MySQL.DB.Delete(&user)

		return nil, uid
	}

	return err, user.Uid
}

func UpdatePassword(uid string, srcPwd string, newPwd string) error {

	err, user := Find(uid, srcPwd)

	if err == nil {
		user.Pwd = newPwd

		db.MySQL.DB.Save(&user)

		return nil
	}

	return err
}
