package img

import (
	"errors"

	"github.com/pwcong/img-hosting/db"
	Init "github.com/pwcong/img-hosting/init"
	"github.com/pwcong/img-hosting/model"
	fileutils "github.com/pwcong/img-hosting/utils/file"
	timeutils "github.com/pwcong/img-hosting/utils/time"
	"github.com/pwcong/img-hosting/utils/uuid"

	ImgTableService "github.com/pwcong/img-hosting/service/imgtable"
)

func SaveImage(uid string, filename string, contentType string, data []byte) (error, string) {

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

	tableName := "imgs_" + year + "_" + month

	if !db.MySQL.DB.HasTable(tableName) {
		db.MySQL.DB.CreateTable(&model.Img{TBName: tableName})
		ImgTableService.AddImageTable(tableName)
	}

	img := &model.Img{
		Filename:  filename,
		Type:      contentType,
		Data:      data,
		Year:      year,
		Month:     month,
		Day:       day,
		Storename: storename,
		Path:      path,
		Uid:       uid,
		CreatedAt: timeutils.Now(),

		TBName: tableName,
	}

	db.MySQL.DB.Create(img)

	notFound := db.MySQL.DB.First(&model.Img{TBName: tableName}, "path = ?", img.Path).RecordNotFound()

	if notFound {
		return errors.New("unknown error"), ""
	}

	return nil, img.Path

}

func LoadImage(year string, month string, day string, storename string) (error, model.Img) {

	img := model.Img{
		TBName: "imgs_" + year + "_" + month,
	}

	if !db.MySQL.DB.HasTable(&img) {
		return errors.New("not found"), model.Img{}
	}

	notFound := db.MySQL.DB.First(&img, "day = ? AND storename = ?", day, storename).RecordNotFound()

	if notFound {
		return errors.New("not found"), model.Img{}
	}

	return nil, img

}

func DeleteImage(uid string, year string, month string, day string, storename string) error {

	err, img := LoadImage(year, month, day, storename)

	if err != nil {
		return err
	}

	if uid != img.Uid {
		return errors.New("")
	}

	db.MySQL.DB.Delete(&img)

	return nil
}

func GetImages(uid string) (error, []model.Img) {

	var res []model.Img

	_, imgTables := ImgTableService.GetImageTables()

	for _, imgTable := range imgTables {

		var imgs []model.Img

		db.MySQL.DB.Table(imgTable.TBName).Select("filename, year, month, day, storename, path").Find(&imgs, "uid = ?", uid)

		for _, img := range imgs {
			res = append(res, img)
		}

	}

	return nil, res

}
