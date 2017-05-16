package model

type ImgTable struct {
	TBName string `gorm:"type:varchar(255);primary_key"`
}

func (imgTable ImgTable) TableName() string {

	return "imgs_tables"

}
