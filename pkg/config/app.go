package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	//Open("사용하고자하는 DB", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "sw90lee:sw90lee!@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}