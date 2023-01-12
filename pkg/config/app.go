package config

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:P@sSw0Rd!/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
