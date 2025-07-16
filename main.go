package main

import (
	"net/http"
	"seaals-api/seaals"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/seal", func(c *gin.Context) {
		sm := NewSealMagick()
		// Load the image
		sm.SealMW.ReadImage(seaals.GetRandomSeal())

		so := seaals.ParseQueryOpts(c)
		sm.ApplyEffects(so)

		// Get bytes and return
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	r.GET("/seal/says/:text", func(c *gin.Context) {
		sm := NewSealMagick()
		text := c.Params.ByName("text")

		so := seaals.ParseQueryOpts(c)

		// Load the image
		sm.SealMW.ReadImage(seaals.GetRandomSeal())
		sm.ApplyEffects(so)

		// Draw text
		sm.ThickOutline(text, "Adwaita-Mono", so.Gravity)

		// Get bytes and return
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	return r
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
