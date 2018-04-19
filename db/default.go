package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ORM struct {
	DB   *gorm.DB
	Name string
}

func (o *ORM) Open(user string, password string, address string, dbname string) {

	connectionURL := user + ":" + password + "@" + "tcp(" + address + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(o.Name, connectionURL)

	if err == nil {
		o.DB = db
	} else {
		log.Fatal(err.Error())
	}

}

func (o *ORM) Close() {
	if o.DB != nil {
		o.DB.Close()
	}
}
