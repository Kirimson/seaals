package magick

import (
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type Opts struct {
	Position     string
	Filter       string
	FontSize     float64
	FontColour   string
	BorderColour string
	BorderSize   float64
	// Parsed from user provided Position
	Gravity imagick.GravityType
}

func defaultFloat(value float64, defVal float64) float64 {
	if value == 0 {
		return defVal
	}
	return value
}

func defaultString(value string, defVal string) string {
	if value == "" {
		return defVal
	}
	return value
}

func DefaultOpts(so *Opts) *Opts {
	so.parsePosition()
	so.FontSize = defaultFloat(so.FontSize, 48)
	so.FontColour = defaultString(so.FontColour, "white")
	so.BorderColour = defaultString(so.BorderColour, "black")
	so.BorderSize = defaultFloat(so.BorderSize, 6)
	return so
}

func (so *Opts) parsePosition() {
	if so.Position == "" {
		so.Gravity = imagick.GRAVITY_SOUTH
		return
	}

	if strings.ToLower(so.Position) == "top" {
		so.Gravity = imagick.GRAVITY_NORTH
	} else if strings.ToLower(so.Position) == "bottom" {
		so.Gravity = imagick.GRAVITY_SOUTH
	} else if strings.ToLower(so.Position) == "center" {
		so.Gravity = imagick.GRAVITY_CENTER
	} else if strings.ToLower(so.Position) == "left" {
		so.Gravity = imagick.GRAVITY_WEST
	} else if strings.ToLower(so.Position) == "right" {
		so.Gravity = imagick.GRAVITY_EAST
	}
}
