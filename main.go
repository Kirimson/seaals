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
		sm.SealMW.ReadImage(seaals.GetRandomSeal())
		sm.opts = seaals.ParseQueryOpts(c)
		sm.ApplyEffects()
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	r.GET("/seal/says/:text", func(c *gin.Context) {
		sm := NewSealMagick()
		text := c.Params.ByName("text")
		sm.opts = seaals.ParseQueryOpts(c)
		sm.SealMW.ReadImage(seaals.GetRandomSeal())
		sm.ApplyEffects()
		sm.DrawText(text)
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	return r
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	r := setupRouter()
	r.Run(":8080")
}
