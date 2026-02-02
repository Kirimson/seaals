package image

import (
	"errors"
	"fmt"
	"log"
	"seaals/vips"
	"strings"
)

type Image struct {
	Image     *vips.Image
	ImageType vips.ImageType
	Pages     int
	Height    int
	Width     int
}

// LoadImage will load an image from a set path, returning a slice of vips.Image,
// with each index being a page. Still images will consist of a single page. Animated
// images will have a page per frame of animation
func LoadImage(path string) (*Image, error) {
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

	opts := vips.DefaultArrayjoinOptions()
	opts.Hspacing = imgs[0].Width()
	opts.Vspacing = imgs[0].Height()
	joined, err := vips.NewArrayjoin(imgs, opts)
	if err != nil {
		return nil, err
	}

	img := Image{
		Image:     joined,
		ImageType: imgs[0].Format(),
		Pages:     imgs[0].Pages(),
		Height:    imgs[0].Height(),
		Width:     imgs[0].Width(),
	}

	return &img, nil
}

func LoadImageBytes(data []byte) (*Image, error) {
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

	opts := vips.DefaultArrayjoinOptions()
	opts.Hspacing = imgs[0].Width()
	opts.Vspacing = imgs[0].Height()
	joined, err := vips.NewArrayjoin(imgs, opts)
	if err != nil {
		return nil, err
	}
	img := Image{
		Image:     joined,
		ImageType: imgs[0].Format(),
		Pages:     imgs[0].Pages(),
		Height:    imgs[0].Height(),
		Width:     imgs[0].Width(),
	}

	return &img, nil
}

func GetImageBytes(img *Image) ([]byte, error) {
	switch img.ImageType {
	case vips.ImageTypeGif:
		saveOpts := vips.DefaultGifsaveBufferOptions()
		saveOpts.PageHeight = img.Height
		return img.Image.GifsaveBuffer(saveOpts)
	case vips.ImageTypePng:
		return img.Image.PngsaveBuffer(vips.DefaultPngsaveBufferOptions())
	case vips.ImageTypeJpeg:
		return img.Image.JpegsaveBuffer(vips.DefaultJpegsaveBufferOptions())
	}
	return nil, errors.New("failed to output image")
}

// DrawText will composite a text image on top of a slice of vips.Image
func DrawText(img *Image, text string, textOpts *Opts) error {
	// Always set the width of the text to the image's width
	// Assume that at least one image (page) is provided
	textOpts.Width = img.Width
	textOpts.Height = img.Height
	// Create the base Text image
	textImg, err := makeText(text, textOpts)
	// Close the text image after this function. Will only be used in this scope
	if err != nil {
		return err
	}
	defer textImg.Close()

	// Replicate text imag based on amount of pages. For non-animated images, this
	// will always just be 1. For animated images, this will be per frame
	if err := textImg.Replicate(1, img.Pages); err != nil {
		return err
	}

	if err := img.Image.Composite2(textImg, vips.BlendModeOver, &vips.Composite2Options{
		X: 0,
		Y: 0,
	}); err != nil {
		log.Fatalf("failed to add text: %s", err)
	}
	return nil
}

func ApplyEffects(img *Image, filter string) error {
	switch strings.ToLower(filter) {
	case "monochrome":
		return FilterMonochrome(img)
	case "invert":
		return FilterInvert(img)
	}
	return nil
}

func FilterMonochrome(img *Image) error {
	if err := img.Image.Colourspace(vips.InterpretationBW, vips.DefaultColourspaceOptions()); err != nil {
		return err
	}
	return nil
}

func FilterInvert(img *Image) error {
	if err := img.Image.ExtractBand(0, &vips.ExtractBandOptions{N: 3}); err != nil {
		return err
	}
	if err := img.Image.Invert(); err != nil {
		return err
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

	// Calculate the y position based on image height
	var yPos int
	switch opts.Position {
	case TextTop:
		yPos = int(float64(opts.Height)*0.05) + opts.Size/2
	case TextMiddle:
		yPos = int(float64(opts.Height) / 2)
	default:
		yPos = int(float64(opts.Height) * 0.95)
	}

	svg := fmt.Sprintf(`
		<svg width="%d" height="%d">
			<text x="%d" y="%d" text-anchor="%s"
				font-family="sans-serif" font-size="%d" font-weight="bold"
				fill="%s" stroke="%s" stroke-width="%d"
			>
				%s
			</text>
		</svg>`, opts.Width, opts.Height, xPos, yPos, opts.Align, opts.Size, opts.Colour, opts.Stroke, opts.StrokeSize, text)
	return vips.NewImageFromBuffer([]byte(svg), nil)
}
