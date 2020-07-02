package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Obj struct {
}

var gdb *gorm.DB

func NewDbCon() *gorm.DB {
	db, err := gorm.Open("mysql", `root:WOaini1314@tcp(test.bybyte.cn:3306)/ny?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`)
	fmt.Println(err)
	return db
}
func GetDB() *gorm.DB {
	if gdb != nil {
		return gdb
	}
	db, err := gorm.Open("mysql", `root:WOaini1314@tcp(test.bybyte.cn:3306)/ny?charset=utf8&parseTime=True&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s`)
	db.LogMode(false)
	gdb = db
	fmt.Println(err)
	return db
}

func BeginTx() {

}
