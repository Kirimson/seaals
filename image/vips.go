package image

import (
	"errors"
	"fmt"
	"log"
	"seaals/vips"
	"strings"
)

// LoadImage will load an image from a set path, returning a slice of vips.Image,
// with each index being a page. Still images will consist of a single page. Animated
// images will have a page per frame of animation
func LoadImage(path string) ([]*vips.Image, error) {
	var imgs []*vips.Image
	meta, err := vips.NewImageFromFile(path, &vips.LoadOptions{})
	if err != nil {
		return nil, err
	}
	pages := meta.Pages()
	for page := range pages {
		img, err := vips.NewImageFromFile(path, &vips.LoadOptions{Page: page})
		if err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func LoadImageBytes(data []byte) ([]*vips.Image, error) {
	var imgs []*vips.Image
	meta, err := vips.NewImageFromBuffer(data, &vips.LoadOptions{})
	if err != nil {
		return nil, err
	}
	pages := meta.Pages()
	for page := range pages {
		img, err := vips.NewImageFromBuffer(data, &vips.LoadOptions{Page: page})
		if err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func GetImageBytes(imgs []*vips.Image) ([]byte, error) {
	opts := vips.DefaultArrayjoinOptions()
	opts.Hspacing = imgs[0].Width()
	opts.Vspacing = imgs[0].Height()
	joined, err := vips.NewArrayjoin(imgs, opts)
	if err != nil {
		return nil, err
	}

	pageHeight := imgs[0].Height()
	switch imgs[0].Format() {
	case vips.ImageTypeGif:
		saveOpts := vips.DefaultGifsaveBufferOptions()
		saveOpts.PageHeight = pageHeight
		return joined.GifsaveBuffer(saveOpts)
	case vips.ImageTypePng:
		return joined.PngsaveBuffer(vips.DefaultPngsaveBufferOptions())
	case vips.ImageTypeJpeg:
		return joined.JpegsaveBuffer(vips.DefaultJpegsaveBufferOptions())
	}
	return nil, errors.New("failed to output image")
}

// DrawText will composite a text image on top of a slice of vips.Image
func DrawText(imgs []*vips.Image, text string, textOpts *Opts) error {
	// Always set the width of the text to the image's width
	// Assume that at least one image (page) is provided
	textOpts.Width = imgs[0].Width()
	// Create the base Text image
	textImg, err := makeText(text, textOpts)
	// Close the text image after this function. Will only be used in this scope
	if err != nil {
		return err
	}
	defer textImg.Close()

	// Calculate the y position based on image height
	var yPos int
	switch textOpts.Position {
	case TextTop:
		yPos = int(float64(imgs[0].Height()) * 0.05)
	case TextMiddle:
		yPos = int(float64(imgs[0].Height())/2 - float64(textImg.Height())/2)
	default:
		yPos = int(float64(imgs[0].Height())*0.95) - textImg.Height()
	}

	// Iterate over all image pages, compositing text image on top
	for _, img := range imgs {
		if err := img.Composite2(textImg, vips.BlendModeOver, &vips.Composite2Options{
			X: 0,
			Y: yPos,
		}); err != nil {
			log.Fatalf("failed to add text: %s", err)
		}
	}
	return nil
}

func ApplyEffects(imgs []*vips.Image, filter string) error {
	switch strings.ToLower(filter) {
	case "monochrome":
		return FilterMonochrome(imgs)
	case "invert":
		return FilterInvert(imgs)
	}
	return nil
}

func FilterMonochrome(imgs []*vips.Image) error {
	for _, img := range imgs {
		if err := img.Colourspace(vips.InterpretationBW, vips.DefaultColourspaceOptions()); err != nil {
			return err
		}
	}
	return nil
}

func FilterInvert(imgs []*vips.Image) error {
	for _, img := range imgs {
		if err := img.ExtractBand(0, &vips.ExtractBandOptions{N: 3}); err != nil {
			return err
		}
		if err := img.Invert(); err != nil {
			return err
		}
	}
	return nil
}

func makeText(text string, opts *Opts) (*vips.Image, error) {
	// Set xPos based on alignment
	var xPos int
	switch opts.Align {
	case TextLeft:
		xPos = 0
	case TextRight:
		xPos = opts.Width
	default:
		xPos = opts.Width / 2
	}
	svg := fmt.Sprintf(`
		<svg width="%d" height="%d">
			<text x="%d" y="%d" text-anchor="%s"
				font-family="sans-serif" font-size="%d" font-weight="bold"
				fill="%s" stroke="%s" stroke-width="%d"
			>
				%s
			</text>
		</svg>`, opts.Width, opts.Size, xPos, opts.Size-(opts.StrokeSize*2), opts.Align, opts.Size, opts.Colour, opts.Stroke, opts.StrokeSize, text)
	return vips.NewImageFromBuffer([]byte(svg), nil)
}
