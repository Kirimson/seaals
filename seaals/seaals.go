package seaals

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type SealOpts struct {
	Position string `form:"position" json:"position"`
	Filter   string `form:"filter" json:"filter"`
	// Parsed from user provided Position
	Gravity imagick.GravityType
}

func ParseQueryOpts(c *gin.Context) *SealOpts {
	var so SealOpts
	c.ShouldBind(&so)
	so.parsePosition()
	return &so
}

func (so *SealOpts) parsePosition() {
	if so.Position == "" {
		so.Gravity = imagick.GRAVITY_SOUTH
		return
	}

	if strings.ToLower(so.Position) == "top" {
		so.Gravity = imagick.GRAVITY_NORTH
	} else if strings.ToLower(so.Position) == "bottom" {
		so.Gravity = imagick.GRAVITY_SOUTH
	}
}

func GetRandomSeal() string {
	// TODO: Actually make it do something
	return "seal.jpeg"
}
