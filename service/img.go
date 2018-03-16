package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pwcong/img-hosting/model"
)

type ImgService struct {
	Base *BaseService
}

func GetExtension(filename string) string {

	if len(filename) < 1 {
		return "unknown"
	}

	res := strings.Split(filename, ".")

	if len(res) < 2 {
		return "unknown"
	}

	return res[1]
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
			ExtName:  GetExtension(file.Filename),
		}

		dir := filepath.Join(filepath.Dir(os.Args[0]), "public/"+img.Year+"/"+img.Month+"/"+img.Date)

		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return model.Img{}, err
		}

		err = ioutil.WriteFile(filepath.Join(dir, img.Symbol+"."+img.ExtName), buffer, 0666)
		if err != nil {
			return model.Img{}, err
		}

		db.Create(&img)
	}

	return img, nil
}

func (ctx *ImgService) SavePrivateImage(file *multipart.FileHeader, id int) (model.Img, error) {

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

	db.Model(&user).Association("Imgs").Append([]model.Img{img})

	return img, nil

}
