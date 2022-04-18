package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"goginproj/Infra"
	"goginproj/dtos"
	"goginproj/models"
	"goginproj/services"
)

func RegisterBookRoutes(route *gin.RouterGroup) {
	route.GET("/", FindBooks)
	route.POST("/", CreateBook)
	route.PATCH("/", UpdateBook)
	route.GET("/:id", FindBook)
}

func FindBooks(c *gin.Context) {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	log.Debugf("Query page size %v with page %v", pageSize, page)
	books, count := services.GetAllBooks(page, pageSize)
	c.JSON(http.StatusOK, dtos.CreatedBookPagedResponse(books, page, pageSize, count))
}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := Infra.GetDB().Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := middlewares.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrNotFound.Message})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	middlewares.DB.Model(&book).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": book})

}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	middlewares.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
