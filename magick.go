// Port of http://members.shaw.ca/el.supremo/MagickWand/trans_paint.htm to Go
package main

import (
	"regexp"
	"seaals-api/seaals"
	"strings"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type SealMagick struct {
	// MagickWand instance for the main Seal image
	// It holds the actual image data. Multiple MagickWands
	// can be used during operations, but should always be
	// ultimately applied to the SealMW
	SealMW *imagick.MagickWand
	// The DrawingWand is the 'tool' to add things to
	// a MagickWand's image
	Dw      *imagick.DrawingWand
	Details ImageDetails
	opts    *seaals.SealOpts
	aw      *imagick.MagickWand
}

type ImageDetails struct {
	MimeType string
}

func NewSealMagick() *SealMagick {
	return &SealMagick{
		SealMW: imagick.NewMagickWand(),
		Dw:     imagick.NewDrawingWand(),
		aw:     imagick.NewMagickWand(),
	}
}

func (sm *SealMagick) DrawText(text string) {
	defer sm.Dw.Clear()
	// PixelWands are the properties of a DrawingWand
	// controlling the colour/alpha of what is drawn
	textPW := imagick.NewPixelWand()

	// Set up a 72 point white font
	textPW.SetColor(sm.opts.FontColor)
	sm.Dw.SetFillColor(textPW)
	sm.Dw.SetFont("Adwaita-Mono")
	sm.Dw.SetFontSize(sm.opts.FontSize)

	// Add a black outline to the text
	textPW.SetColor(sm.opts.BorderColor)
	sm.Dw.SetStrokeColor(textPW)
	sm.Dw.SetStrokeWidth(sm.opts.BorderStroke)

	sm.Dw.SetGravity(sm.opts.Gravity)
	sm.Dw.Annotation(0, 0, text)

	// Redraw the text, with border colour set to none to 'remove' the inside border
	textPW.SetColor("none")
	sm.Dw.SetStrokeColor(textPW)
	sm.Dw.Annotation(0, 0, text)

	for i := 0; i < int(sm.SealMW.GetNumberImages()); i++ {
		sm.SealMW.SetIteratorIndex(i)
		sm.SealMW.DrawImage(sm.Dw)
	}
}

func (sm *SealMagick) ApplyEffects() {
	if strings.ToLower(sm.opts.Filter) == "monochrome" {
		sm.FilterMonochrome()
	} else if strings.ToLower(sm.opts.Filter) == "funky" {
		sm.FilterFunky()
	} else if strings.ToLower(sm.opts.Filter) == "invert" {
		sm.FilterInvert()
	}
}

func (sm *SealMagick) FilterMonochrome() {
	for i := 0; i < int(sm.SealMW.GetNumberImages()); i++ {
		sm.SealMW.SetIteratorIndex(i)
		sm.SealMW.TransformImageColorspace(imagick.COLORSPACE_GRAY)
	}
}

func (sm *SealMagick) FilterInvert() {
	for i := 0; i < int(sm.SealMW.GetNumberImages()); i++ {
		sm.SealMW.SetIteratorIndex(i)
		sm.SealMW.NegateImage(false)
	}
}

func (sm *SealMagick) FilterFunky() {
	// Need a monochrome image to apply colour filters
	sm.FilterMonochrome()

	gradientMW := imagick.NewMagickWand()
	gradientMW.SetSize(300, 30)

	fc := imagick.NewPixelWand()
	fc.SetColor("none")
	gradientMW.NewImage(300, 30, fc)
	gradientMW.SetImageArtifact("gradient:angle", "90")

	red := imagick.NewPixelWand()
	if ok := red.SetColor("red"); !ok {
		panic(gradientMW.GetLastError())
	}

	green := imagick.NewPixelWand()
	if ok := green.SetColor("green"); !ok {
		panic(gradientMW.GetLastError())
	}

	blue := imagick.NewPixelWand()
	if ok := blue.SetColor("blue"); !ok {
		panic(gradientMW.GetLastError())
	}

	// Make the gradient with the PixelWand colours
	stops := []imagick.StopInfo{
		imagick.NewStopInfo(red.GetMagickColor(), 0),
		imagick.NewStopInfo(green.GetMagickColor(), 0.50),
		imagick.NewStopInfo(blue.GetMagickColor(), 1),
	}
	err := gradientMW.GradientImage(imagick.GRADIENT_TYPE_LINEAR, imagick.SPREAD_METHOD_PAD, stops)
	if err != nil {
		panic(err)
	}

	// Apply the CLUT to the base image
	for i := 0; i < int(sm.SealMW.GetNumberImages()); i++ {
		sm.SealMW.SetIteratorIndex(i)
		sm.SealMW.ClutImage(gradientMW, imagick.INTERPOLATE_PIXEL_AVERAGE)
	}
}

func (sm *SealMagick) LoadImage(image string) {
	sm.SealMW.ReadImage(image)
	sm.Details = sm.IdentifyImage()
}

func (sm *SealMagick) IdentifyImage() ImageDetails {
	identifyString := sm.SealMW.IdentifyImage()
	mimeRegexp, _ := regexp.Compile("Mime type: ([a-z]+/([a-z]+))")
	mimeMatch := mimeRegexp.FindStringSubmatch(identifyString)

	mime := ""
	if len(mimeMatch) == 3 {
		mime = mimeMatch[1]
	}
	return ImageDetails{
		MimeType: mime,
	}
}

func (sm *SealMagick) GetImageBytes() ([]byte, error) {
	return sm.SealMW.GetImagesBlob()
}
