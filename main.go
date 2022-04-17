package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"goginproj/controllers"
	"goginproj/models"
)

func main() {
	r := gin.Default()
	log.Info("starting ...")

	models.ConnectDB()
	r.GET("/book", controllers.FindBooks)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)

	r.POST("/create", controllers.CreatBook)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": 200,
		})
	})

	r.Run(":3000")
}
