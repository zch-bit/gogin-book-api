package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(fmt.Sprintln("Failed to connect with databse", err.Error()))
	}
	database.AutoMigrate(&Book{})

	DB = database
}
