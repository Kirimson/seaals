package controller

import (
	"fmt"
	"path/filepath"
	"seaals/magick"
	"seaals/models"
	"seaals/service"
)

// SealController interacts with the Magick package, to modify Seal images with
// different effects
type SealController struct {
	sealService *service.SealService
	basePath    string
}

type SealResponse struct {
	Image    []byte
	MimeType string
}

// Return a new SealController, that will call methods from the provided SealService
// SealController will implement the logic for different Seal-related routes
func NewSealController(service *service.SealService, basePath string) *SealController {
	return &SealController{
		sealService: service,
		basePath:    basePath,
	}
}

// GetSeal gets a random Seal image, applies filters and
// returns a SealResponse containing the image bytes and MimeType
func (sc *SealController) GetSeal(seal *models.Seal, mo *magick.Opts) (*SealResponse, error) {
	sm := magick.NewSealMagick(mo)
	sm.LoadImage(seal.Path)

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
func (sc *SealController) GetSealSaying(seal *models.Seal, text string, mo *magick.Opts) (*SealResponse, error) {
	sm := magick.NewSealMagick(mo)
	sm.LoadImage(seal.Path)
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

func (sc *SealController) RandomSeal() models.Seal {
	// TODO: Actually make it do something
	path := filepath.Join(sc.basePath, "seal.jpeg")
	fmt.Println(path)
	return models.Seal{
		Path:     path,
		MimeType: "image/jpeg",
		Tags:     []models.Tag{{Name: "cute"}},
	}
}
