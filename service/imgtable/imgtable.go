package imgtable

import (
	"errors"

	"github.com/pwcong/img-hosting/db"
	"github.com/pwcong/img-hosting/model"
)

func AddImageTable(tableName string) error {

	db.MySQL.DB.Create(model.ImgTable{TBName: tableName})

	notFound := db.MySQL.DB.First(&model.ImgTable{}, "tb_name = ?", tableName).RecordNotFound()

	if notFound {
		return errors.New("unknown error")
	}

	return nil

}

func GetImageTables() (error, []model.ImgTable) {

	var imgTables []model.ImgTable

	db.MySQL.DB.Find(&imgTables)

	return nil, imgTables

}
