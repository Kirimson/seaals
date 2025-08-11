/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"seaals/api"
	"seaals/api/middleware"
	"seaals/controller"
	"seaals/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/urfave/cli/v3"
	"gopkg.in/gographics/imagick.v3/imagick"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

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

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return s
}

func Serve(ctx context.Context, cmd *cli.Command) error {
	database := cmd.String("database")
	port := strconv.Itoa(int(cmd.Int16("port")))

	imagick.Initialize()
	defer imagick.Terminate()

	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Service interfaces with the database
	sealService := service.NewSealService(db)
	// Controller implements routes, and calls the service
	sealController := controller.NewSealController(sealService)

	// Create an instance of the API server
	seaalsApi := api.NewSeaalsServer(sealController)

	s := newRouter(seaalsApi, port)
	// Run the HTTP server, and return any error it returns
	return s.ListenAndServe()
}
