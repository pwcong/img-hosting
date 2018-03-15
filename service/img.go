package service

import (
	"crypto/md5"
	"encoding/hex"
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

func (ctx *ImgService) SaveImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err

	}

	defer src.Close()

	buffer, err := ioutil.ReadAll(src)
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = h.Write(buffer)
	if err != nil {
		return "", err
	}

	id := hex.EncodeToString(h.Sum(nil))

	var img model.Img

	db := ctx.Base.DB
	now := time.Now()

	notFound := db.Where("id = ?", id).First(&img).RecordNotFound()
	if notFound {
		img = model.Img{
			ID:        id,
			Filename:  file.Filename,
			Year:      now.Format("2006"),
			Month:     now.Format("01"),
			Date:      now.Format("02"),
			ExtName:   GetExtension(file.Filename),
			CreatedAt: now,
			UpdatedAt: now,
		}

		dir := filepath.Join(filepath.Dir(os.Args[0]), "public/"+img.Year+"/"+img.Month+"/"+img.Date)

		err = os.MkdirAll(dir, 0666)
		if err != nil {
			return "", err
		}

		err = ioutil.WriteFile(filepath.Join(dir, img.ID+"."+img.ExtName), buffer, 0666)
		if err != nil {
			return "", err
		}

		db.Create(img)
	}

	return "/public/" + img.Year + "/" + img.Month + "/" + img.Date + "/" + img.ID + "." + img.ExtName, nil
}
