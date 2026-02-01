package controller

import (
	"path/filepath"
	"seaals/image"
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

// GetSealImage gets a random Seal image, applies filters and
// returns a SealResponse containing the image bytes and MimeType
func (sc *SealController) GetSealImage(seal *models.Seal, opts *image.Opts) (*SealResponse, error) {
	imgs, err := image.LoadImage(seal.Path)
	if err != nil {
		return nil, err
	}

	if err := image.ApplyEffects(imgs, opts.Filter); err != nil {
		return nil, err
	}
	b, err := image.GetImageBytes(imgs)
	if err != nil {
		return nil, err
	}

	resp := &SealResponse{
		Image:    b,
		MimeType: string(imgs[0].Format()),
	}
	return resp, nil
}

// GetSealSaying gets a random Seal with both graphical effects and a caption text
func (sc *SealController) GetSealSaying(seal *models.Seal, text string, opts *image.Opts) (*SealResponse, error) {
	imgs, err := image.LoadImage(seal.Path)
	if err != nil {
		return nil, err
	}
	if err := image.ApplyEffects(imgs, opts.Filter); err != nil {
		return nil, err
	}

	// Draw text after all effects have been applied
	if err := image.DrawText(imgs, text, opts); err != nil {
		return nil, err
	}

	b, err := image.GetImageBytes(imgs)
	if err != nil {
		return nil, err
	}

	resp := &SealResponse{
		Image:    b,
		MimeType: string(imgs[0].Format()),
	}
	return resp, nil
}

func (sc *SealController) GetSealByPath(path string) (*models.Seal, error) {
	seal, err := sc.sealService.GetSealByPath(path)
	if err != nil {
		return nil, err
	}
	// Make a real path from the relative path from basePath
	seal.Path = filepath.Join(sc.basePath, seal.Path)
	return seal, nil
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

func (sc *SealController) CountSeals() (int64, error) {
	return sc.sealService.CountSeals()
}

func (sc *SealController) GetPopularTags() ([]*models.TagStat, error) {
	return sc.sealService.GetPopularTags()
}
