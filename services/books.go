package services

import (
	"github.com/jinzhu/gorm"
	"goginproj/Infra"
	"goginproj/models"
)

func init() {

}

type Repository struct {
	DB *gorm.DB
}

func GetAllBooks(page, pageSize int) ([]models.Book, int) {
	var books []models.Book
	var count int

	db := Infra.GetDB()
	db.Model(&books).Count(&count)
	db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&books)
	return books, count
}
