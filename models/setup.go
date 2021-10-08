package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("gagal konection")
	}

	db.AutoMigrate(&Mahasiswa{})

	return db
}
