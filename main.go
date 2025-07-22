package main

import (
	"net/http"
	"seaals-api/seaals"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/seal", func(c *gin.Context) {
		sm := NewSealMagick()
		sm.LoadImage(seaals.GetRandomSeal())
		sm.opts = seaals.ParseQueryOpts(c)
		sm.ApplyEffects()
		b, err := sm.GetImageBytes()
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Data(http.StatusOK, sm.Details.MimeType, b)
	})

	r.GET("/seal/says/:text", func(c *gin.Context) {
		sm := NewSealMagick()
		text := c.Params.ByName("text")
		sm.opts = seaals.ParseQueryOpts(c)
		sm.LoadImage(seaals.GetRandomSeal())
		sm.ApplyEffects()
		sm.DrawText(text)
		b, err := sm.GetImageBytes()
		if err != nil {
			c.AbortWithError(500, err)
		}
		c.Data(http.StatusOK, sm.Details.MimeType, b)
	})

	return r
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	r := setupRouter()
	r.Run(":8080")
}
