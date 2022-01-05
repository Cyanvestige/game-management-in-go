package config

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	// replace this part with your db information
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}