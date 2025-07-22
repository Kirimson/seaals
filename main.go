package main

import (
	"seaals-api/controller"
	"seaals-api/models"
	"seaals-api/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schemas
	db.AutoMigrate(&models.Seal{}, &models.Tag{})

	// Service interfaces with the database
	sealService := service.NewSealService(db)
	// Controller implements routes, and calls the service
	sealController := controller.NewSealController(sealService)

	r.GET("/seal", sealController.GetSeal)
	r.GET("/seal/says/:text", sealController.GetSealSaying)

	return r
}

func main() {
	// Setup ImageMagick
	imagick.Initialize()
	defer imagick.Terminate()

	r := setupRouter()
	r.Run(":8080")
}
