package initDB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
//
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/zero_distance")

	if err != nil {
		log.Panicln("err from mysql connect:", err.Error())
	}
}
