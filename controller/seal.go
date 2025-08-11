package controller

import (
	"seaals/magick"
	"seaals/service"
)

// SealController interacts with the Magick package, to modify Seal images with
// different effects
type SealController struct {
	sealService *service.SealService
}

type SealResponse struct {
	Image    []byte
	MimeType string
}

// Return a new SealController, that will call methods from the provided SealService
// SealController will implement the logic for different Seal-related routes
func NewSealController(service *service.SealService) *SealController {
	return &SealController{
		sealService: service,
	}
}

// GetSeal gets a random Seal image, applies filters and
// returns a SealResponse containing the image bytes and MimeType
func (sc *SealController) GetSeal(mo *magick.Opts) (*SealResponse, error) {
	sm := magick.NewSealMagick(mo)
	sm.LoadImage(sc.randomSeal())

	sm.ApplyEffects()
	b, err := sm.GetImageBytes()
	if err != nil {
		return nil, err
	}

	resp := &SealResponse{
		Image:    b,
		MimeType: sm.Details.MimeType,
	}
	return resp, nil
}

// GetSealSaying gets a random Seal with both graphical effects and a caption text
func (sc *SealController) GetSealSaying(text string, mo *magick.Opts) (*SealResponse, error) {
	sm := magick.NewSealMagick(mo)
	sm.LoadImage(sc.randomSeal())

	sm.LoadImage(sc.randomSeal())
	sm.ApplyEffects()

	// Draw text after all effects have been applied
	sm.DrawText(text)

	b, err := sm.GetImageBytes()
	if err != nil {
		return nil, err
	}

	resp := &SealResponse{
		Image:    b,
		MimeType: sm.Details.MimeType,
	}
	return resp, nil
}

func (sc *SealController) randomSeal() string {
	// TODO: Actually make it do something
	return "seal.jpeg"
}
