/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"net"
	"net/http"
	"seaals/api"
	"seaals/api/middleware"
	"seaals/controller"
	"seaals/models"
	"seaals/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func newRouter(seaals *api.Server, port string) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("error loading swagger spec\n %s", err)
	}
	swagger.Servers = nil

	r := gin.Default()
	// r.Use(oapimiddleware.OapiRequestValidator(swagger))
	r.Use(middleware.ErrorHandler())

	api.RegisterSwagger(r)

	api.RegisterHandlers(r, seaals)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Static("/assets", "./public/assets")
	r.LoadHTMLGlob("public/html/*.html")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return s
}

func Serve(ctx context.Context, cmd *cli.Command) error {
	database := cmd.String("database")
	basePath := cmd.String("base-path")
	port := strconv.Itoa(int(cmd.Int16("port")))

	imagick.Initialize()
	defer imagick.Terminate()

	sealDB, err := models.InitialiseDB(database)
	if err != nil {
		return err
	}

	// Service interfaces with the database
	sealService := service.NewSealService(sealDB)
	// Controller implements core business logic for SEAaLS
	sealController := controller.NewSealController(sealService, basePath)

	// Create an instance of the API server which implements routes
	seaalsApi := api.NewSeaalsServer(sealController)

	s := newRouter(seaalsApi, port)
	// Run the HTTP server, and return any error it returns
	return s.ListenAndServe()
}
