package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	conn := "root:$0w$K@Ny7kM6@tcp(localhost:3306)/sakila"

	db, err := gorm.Open(
		mysql.Open(conn+"?parseTime=true"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
	DB = db

}
