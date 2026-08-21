/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"seaals/api"
	"seaals/controller"
	"seaals/models"
	"seaals/service"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/urfave/cli/v3"
)

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func newRouter(seaals *api.Server, port string) *http.Server {
	r := chi.NewMux()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Add generate Api server implementation to Chi Mux
	h := api.HandlerFromMux(seaals, r)

	// Register Swagger
	api.RegisterSwagger(r)

	// Serve index page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/html/index.html")
	})

	// Serve static files
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public/assets"))
	FileServer(r, "/assets", filesDir)

	// 404
	// r.NotFound(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "public/html/404.html")
	// })

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return s
}

func Serve(ctx context.Context, cmd *cli.Command) error {
	database := cmd.String("database")
	basePath := cmd.String("base-path")
	port := strconv.Itoa(int(cmd.Int16("port")))

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

	// Create a http.Server using the API server, and some middleware
	s := newRouter(seaalsApi, port)
	// Run the HTTP server, and return any error it returns
	return s.ListenAndServe()
}
