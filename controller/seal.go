package controller

import (
	"net/http"
	"seaals-api/magick"
	"seaals-api/service"

	"github.com/gin-gonic/gin"
)

// SealController will implement buisness logic for different routes
// This is usually a 'service', and our service is normall a
// 'repository' but I feel there would just be too much duplication that way
// SealController interacts with the Magick package, to modify Seal images with
// different effects
type SealController struct {
	sealService *service.SealService
}

// Return a new SealController, that will call methods from the provided SealService
// SealController will implement the logic for different Seal-related routes
func NewSealController(service *service.SealService) *SealController {
	return &SealController{
		sealService: service,
	}
}

func (sc *SealController) GetSeal(c *gin.Context) {
	var so magick.SealOpts
	c.ShouldBind(&so)

	sm := magick.NewSealMagick(so)
	sm.LoadImage(sc.randomSeal())

	sm.ApplyEffects()
	b, err := sm.GetImageBytes()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.Data(http.StatusOK, sm.Details.MimeType, b)
}

func (sc *SealController) GetSealSaying(c *gin.Context) {
	var so magick.SealOpts
	c.ShouldBind(&so)

	sm := magick.NewSealMagick(so)
	sm.LoadImage(sc.randomSeal())

	sm.LoadImage(sc.randomSeal())
	sm.ApplyEffects()

	text := c.Params.ByName("text")
	sm.DrawText(text)

	b, err := sm.GetImageBytes()
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.Data(http.StatusOK, sm.Details.MimeType, b)
}

func (sc *SealController) randomSeal() string {
	// TODO: Actually make it do something
	return "snow.gif"
}
