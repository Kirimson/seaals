package controller

import (
	"fmt"
	"os"
	"path/filepath"
	"seaals/models"
	"seaals/service"
	"slices"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

type SealLI struct {
	sealService *service.SealService
	basePath    string
}

var AllowedMimes = []string{"image/jpeg", "image/gif", "image/png"}

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

func (sli *SealLI) AddSeal(sealData []byte, tags []string) (*models.Seal, error) {
	mtype := mimetype.Detect(sealData)
	if !slices.Contains(AllowedMimes, mtype.String()) {
		return nil, fmt.Errorf("file format %s is not allowed", mtype.String())
	}

	// Add the seal file to the current base directory
	sealUuid := uuid.New()
	fileName := fmt.Sprintf("%s%s", strings.ReplaceAll(sealUuid.String(), "-", ""), mtype.Extension())
	os.WriteFile(filepath.Join(sli.basePath, fileName), sealData, 0660)

	// Ensure all tags exist for this new Seal
	var sealTags []models.Tag
	for _, tagName := range tags {
		tag, err := sli.sealService.GetTagByName(tagName)
		if err != nil {
			return nil, err
		}
		// If tag does not exist, create it
		if tag == nil {
			tag, err = sli.AddTag(tagName)
			if err != nil {
				return nil, err
			}
		}
		// Add the retrieved or newly made tag to the []Tag
		sealTags = append(sealTags, *tag)
	}

	seal := &models.Seal{
		Path:     fileName,
		Tags:     sealTags,
		MimeType: mtype.String(),
	}
	seal, err := sli.sealService.CreateSeal(seal)
	if err != nil {
		return nil, err
	}
	return seal, nil
}

func (sli *SealLI) GetAllTags() ([]models.Tag, error) {
	return sli.sealService.GetAllTags()
}

func (sli *SealLI) AddTag(name string) (*models.Tag, error) {
	tag := &models.Tag{
		Name: name,
	}
	tag, err := sli.sealService.CreateTag(*tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
