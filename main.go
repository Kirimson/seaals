package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"seaals-api/api"
	"seaals-api/controller"
	"seaals-api/middleware"
	"seaals-api/models"
	"seaals-api/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	oapimiddleware "github.com/oapi-codegen/gin-middleware"
)

func newRouter(seaals *api.Server, port string) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("error loading swagger spec\n %s", err)
	}
	swagger.Servers = nil

	r := gin.Default()
	r.Use(oapimiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.ErrorHandler())

	api.RegisterHandlers(r, seaals)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return s
}

func main() {
	// Setup ImageMagick
	imagick.Initialize()
	defer imagick.Terminate()

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

	port := flag.String("port", "8080", "Port for the HTTP server")
	flag.Parse()
	// Create an instance of the API server
	seaalsApi := api.NewSeaalsServer(sealController)

	s := newRouter(seaalsApi, *port)
	log.Fatal(s.ListenAndServe())
}
