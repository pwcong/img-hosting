package img

import (
	"errors"

	"github.com/pwcong/img-hosting/db"
	Init "github.com/pwcong/img-hosting/init"
	"github.com/pwcong/img-hosting/model"
	fileutils "github.com/pwcong/img-hosting/utils/file"
	timeutils "github.com/pwcong/img-hosting/utils/time"
	"github.com/pwcong/img-hosting/utils/uuid"
)

func SaveImage(uid string, filename string, data []byte) (error, string) {

	ext := fileutils.GetExtension(filename)
	year := timeutils.GetYear()
	month := timeutils.GetMonth()
	day := timeutils.GetDay()

	uuid := uuid.Rand().HexLimit(Init.Config.Server.Image.Limit)

	storename := uuid + "." + ext
	path := "/" + year + "/" + month + "/" + day + "/" + storename

	if uid == "" {
		uid = "guest"
	}

	img := &model.Img{
		Filename:  filename,
		Data:      data,
		Year:      year,
		Month:     month,
		Day:       day,
		Storename: storename,
		Path:      path,
		Uid:       uid,
		CreatedAt: timeutils.Now(),
	}

	db.MySQL.DB.Create(img)

	notFound := db.MySQL.DB.First(&model.Img{}, "path = ?", img.Path).RecordNotFound()

	if notFound {
		return errors.New(""), ""
	}

	return nil, img.Path

}

func LoadImage(year string, month string, day string, storename string) (error, []byte) {

	var img model.Img

	notFound := db.MySQL.DB.First(&img, "year = ? AND month = ? AND day = ? AND storename = ?", year, month, day, storename).RecordNotFound()

	if notFound {
		return errors.New(""), []byte{}

	}

	return nil, img.Data

}
