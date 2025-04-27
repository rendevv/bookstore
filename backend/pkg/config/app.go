package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@Rendy278@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
