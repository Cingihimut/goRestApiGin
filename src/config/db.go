package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	host := "127.0.0.1"
	port := "3306"
	dbName := "rest_api_gin"
	userName := "root"
	password := ""

	dsn := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Asia%2FJakarta"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat terhubung dengan data base")
	}

	return db

}
