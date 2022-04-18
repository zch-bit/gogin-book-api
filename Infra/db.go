package Infra

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(fmt.Sprintln("Failed to connect with database", err.Error()))
	}
	DB = database
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
