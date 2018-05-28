package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/pwcong/img-hosting/model"
	"github.com/pwcong/img-hosting/utils"
)

type ImgService struct {
	Base *BaseService
}

func (ctx *ImgService) SaveImage(file *multipart.FileHeader) (model.Img, error) {
	src, err := file.Open()
	if err != nil {
		return model.Img{}, err
	}

	defer src.Close()

	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return model.Img{}, err
	}

	h := md5.New()
	_, err = h.Write(buffer)
	if err != nil {
		return model.Img{}, err
	}

	symbol := hex.EncodeToString(h.Sum(nil))

	var img model.Img

	db := ctx.Base.DB
	now := time.Now()

	notFound := db.Where("symbol = ?", symbol).First(&img).RecordNotFound()
	if notFound {
		img = model.Img{
			Symbol:   symbol,
			Filename: file.Filename,
			Year:     now.Format("2006"),
			Month:    now.Format("01"),
			Date:     now.Format("02"),
			ExtName:  utils.GetExtension(file.Filename),
		}

		rootDir := utils.GetRootDir()

		dir := filepath.Join(rootDir, "public/"+img.Year+"/"+img.Month+"/"+img.Date)

		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return model.Img{}, err
		}

		err = ioutil.WriteFile(filepath.Join(dir, img.Symbol+"."+img.ExtName), buffer, 0666)
		if err != nil {
			return model.Img{}, err
		}

		if dbc := db.Create(&img); dbc.Error != nil {
			return model.Img{}, dbc.Error
		}
	}

	return img, nil
}

func (ctx *ImgService) SavePrivateImage(file *multipart.FileHeader, id uint) (model.Img, error) {

	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("id = ?", id).First(&user).RecordNotFound()
	if notFound {
		return model.Img{}, errors.New("user is not existed")
	}

	img, err := ctx.SaveImage(file)

	if err != nil {
		return model.Img{}, err
	}

	if dbc := db.Model(&user).Association("Imgs").Append([]model.Img{img}); dbc.Error != nil {
		return model.Img{}, dbc.Error
	}

	return img, nil

}

func (ctx *ImgService) GetPrivateImages(id uint) ([]model.Img, error) {

	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("id = ?", id).First(&user).RecordNotFound()
	if notFound {
		return []model.Img{}, errors.New("user is not existed")
	}

	var imgs []model.Img
	if dbc := db.Model(&user).Related(&imgs, "Imgs"); dbc.Error != nil {
		return []model.Img{}, dbc.Error
	}

	return imgs, nil

}

func (ctx *ImgService) RemovePrivateImages(userId uint, imgIds []uint) ([]model.Img, error) {

	db := ctx.Base.DB

	var user model.User

	notFound := db.Where("id = ?", userId).First(&user).RecordNotFound()
	if notFound {
		return []model.Img{}, errors.New("user is not existed")
	}

	var imgs []model.Img

	db.Where("id in (?)", imgIds).Find(&imgs)

	if dbc := db.Model(&user).Association("Imgs").Delete(imgs); dbc.Error != nil {
		return []model.Img{}, dbc.Error
	}

	return imgs, nil

}
