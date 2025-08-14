package controller

import (
	"seaals/models"
	"seaals/service"
)

type SealLI struct {
	sealService *service.SealService
	basePath    string
}

// Return a new AdminController, that will call methods from the provided SealService
// AdminController will implement admin-related oprtations for SEAaLS
func NewSealLI(service *service.SealService, basePath string) *SealLI {
	return &SealLI{
		sealService: service,
		basePath:    basePath,
	}
}

func (sli *SealLI) GetAllSeals() ([]models.Seal, error) {
	return sli.sealService.GetAllSeals()
}

func (sli *SealLI) AddSeal(path string, tags []string) (*models.Seal, error) {
	// Ensure all tags exist for this new Seal
	var sealTags []models.Tag
	for _, tagName := range tags {
		tag, err := sli.sealService.GetTagByName(tagName)
		if err != nil {
			return nil, err
		}
		// If tag does not exist, create it
		if tag == nil {
			tag = &models.Tag{
				Name: tagName,
			}
			tag, err = sli.sealService.CreateTag(*tag)
			if err != nil {
				return nil, err
			}
		}
		// Add the retrieved or newly made tag to the []Tag
		sealTags = append(sealTags, *tag)
	}

	seal := &models.Seal{
		Path: path,
		Tags: sealTags,
	}
	seal, err := sli.sealService.CreateSeal(seal)
	if err != nil {
		return nil, err
	}
	return seal, nil
}
