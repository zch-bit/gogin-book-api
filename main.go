package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"goginproj/Infra"
	"goginproj/controllers"
	"goginproj/models"
)

type Server struct {
	store  *gorm.DB
	router *gin.Engine
}

func NewServer(store *gorm.DB) *Server {
	goGonicEngine := gin.Default()
	apiRouteGroup := goGonicEngine.Group("/api")
	controllers.RegisterBookRoutes(apiRouteGroup.Group("books"))
	return &Server{
		store:  store,
		router: goGonicEngine,
	}
}

func main() {
	log.Info("Starting...")
	database := Infra.ConnectDatabase()
	defer database.Close()
	database.AutoMigrate(&models.Book{})

	apiServer := NewServer(database)
	apiServer.router.Run(":3000")
}
