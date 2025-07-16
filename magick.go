// Port of http://members.shaw.ca/el.supremo/MagickWand/trans_paint.htm to Go
package main

import (
	"fmt"
	"log"
	"os"
	"seaals-api/seaals"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type SealMagick struct {
	Mw *imagick.MagickWand
	Dw *imagick.DrawingWand
	Pw *imagick.PixelWand
}

func NewSealMagick() *SealMagick {
	return &SealMagick{
		Mw: imagick.NewMagickWand(),
		Dw: imagick.NewDrawingWand(),
		Pw: imagick.NewPixelWand(),
	}
}

func (sm *SealMagick) ThickOutline(text string, font string, position imagick.GravityType) {
	// Set up a 72 point white font
	sm.Pw.SetColor("white")
	sm.Dw.SetFillColor(sm.Pw)
	sm.Dw.SetFont(font)
	sm.Dw.SetFontSize(72)

	// Add a black outline to the text
	sm.Pw.SetColor("black")
	sm.Dw.SetStrokeColor(sm.Pw)
	sm.Dw.SetStrokeWidth(8)

	// Now draw the text, with gravity set to south
	sm.Dw.SetGravity(position)
	sm.Dw.Annotation(0, 65, text)

	// Redraw the text, with border colour set to none to 'remove' the inside border
	sm.Pw.SetColor("none")
	sm.Dw.SetStrokeColor(sm.Pw)
	sm.Dw.Annotation(0, 65, text)
	// Reset gravity
	sm.Dw.SetGravity(imagick.GRAVITY_FORGET)
}

func (sm *SealMagick) ApplyEffects(so *seaals.SealOpts) {
	if strings.ToLower(so.Filter) == "monochrome" {
		sm.Monochrome()
	}
}

func (sm *SealMagick) Monochrome() {
	sm.Mw.SetImageType(imagick.IMAGE_TYPE_GRAYSCALE)
}

// Draw the current stack from DrawingWand to the MagickWand
func (sm *SealMagick) DrawCurrent() {
	sm.Mw.DrawImage(sm.Dw)
}

func (sm *SealMagick) GetImageBytes() []byte {
	// imagick really likes files, likely as we're interfacing with C
	// Make a temp file, read into memory, delete the file, then return
	// TODO: Can this be better?
	id := uuid.New()
	filename := fmt.Sprintf("tmp/%s.jpg", id.String())
	sm.Mw.WriteImage(filename)
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	os.Remove(filename)
	return file
}
