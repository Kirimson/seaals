package controller

import (
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
	if err := sm.LoadImage(seal.Path); err != nil {
		return nil, err
	}

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
	if err := sm.LoadImage(seal.Path); err != nil {
		return nil, err
	}
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

func (sc *SealController) GetSealByID(id int64) (*models.Seal, error) {
	return sc.sealService.GetSealByID(id)
}

func (sc *SealController) RandomSeal(tag *string) (*models.Seal, error) {
	if tag != nil {
		randomSeal, err := sc.sealService.GetRandomSealWithTag(*tag)
		if err != nil {
			return nil, err
		}
		// Make a real path from the relative path from basePath
		randomSeal.Path = filepath.Join(sc.basePath, randomSeal.Path)
		return randomSeal, nil
	}
	randomSeal, err := sc.sealService.GetRandomSeal()
	if err != nil {
		return nil, err
	}
	// Make a real path from the relative path from basePath
	randomSeal.Path = filepath.Join(sc.basePath, randomSeal.Path)
	return randomSeal, nil
}

func (sc *SealController) GetAllSeals() ([]*models.Seal, error) {
	return sc.sealService.GetAllSeals()
}
