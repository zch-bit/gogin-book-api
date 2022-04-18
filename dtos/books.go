package dtos

import (
	"github.com/gin-gonic/gin"
	"goginproj/models"
)

func CreatedBookPagedResponse(books []models.Book, page, pageSize, count int) gin.H {
	totalPage := count / pageSize
	if count%pageSize != 0 {
		totalPage += 1
	}
	return gin.H{
		"books":        books,
		"current_page": page,
		"page_size":    pageSize,
		"total_page":   totalPage,
		"total_books":  count,
	}
}
