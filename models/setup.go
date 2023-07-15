package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cake_store_restful_api"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Cake{})

	DB = database
}
