package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("root:!10-sccrOC/simplerest?charset=utf8&parseTime=True&loc=local"), &gorm.Config{} )
	if err != nil {
		log.Fatal("failed to connect to database: ")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
