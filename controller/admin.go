package controller

import (
	"fmt"
	"os"
	"path/filepath"
	"seaals/magick"
	"seaals/models"
	"seaals/service"
	"slices"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"gopkg.in/gographics/imagick.v3/imagick"
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

func (sli *SealLI) GetAllSeals() ([]*models.Seal, error) {
	seals, err := sli.sealService.GetAllSeals()
	if err != nil {
		return nil, err
	}
	for _, s := range seals {
		tags, err := sli.sealService.GetSealTags(s)
		if err != nil {
			return nil, err
		}
		s.Tags = tags
	}

	return seals, nil
}

func (sli *SealLI) DeleteSeal(id int64) error {
	// Delete the seal file from disk
	seal, err := sli.sealService.GetSealByID(id)
	if err != nil {
		return err
	}
	sealPath := filepath.Join(sli.basePath, seal.Path)
	if err := os.Remove(sealPath); err != nil {
		return fmt.Errorf("failed to delete seal file at %s: %w", sealPath, err)
	}
	// Delete the seal and its tag associations from the database
	return sli.sealService.DeleteSeal(id)
}

func (sli *SealLI) AddSeal(sealData []byte, tags []string) (*models.Seal, error) {
	mtype := mimetype.Detect(sealData)
	if !slices.Contains(AllowedMimes, mtype.String()) {
		return nil, fmt.Errorf("file format %s is not allowed", mtype.String())
	}

	// Add file type as a tag, eg gif, png
	sp := strings.Split(mtype.String(), "/")
	if len(sp) == 2 {
		tags = append(tags, sp[1])
	}

	// Resize the image if it is too big
	imagick.Initialize()
	defer imagick.Terminate()
	sm := magick.NewSealMagick(magick.DefaultOpts(nil))
	sm.LoadImageBytes(sealData)
	sm.ResizeImage(1024)
	sealData, err := sm.GetImageBytes()
	if err != nil {
		return nil, err
	}

	// Add the seal file to the current base directory
	sealUuid := uuid.New()
	fileName := fmt.Sprintf("%s%s", strings.ReplaceAll(sealUuid.String(), "-", ""), mtype.Extension())
	os.WriteFile(filepath.Join(sli.basePath, fileName), sealData, 0660)

	// Ensure all tags exist for this new Seal
	var sealTags []*models.Tag
	for _, tagName := range tags {
		tag, err := sli.sealService.GetTagWithName(tagName)
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
		sealTags = append(sealTags, tag)
	}

	seal := &models.Seal{
		Path:     fileName,
		MimeType: mtype.String(),
		Tags:     sealTags,
	}
	seal, err = sli.sealService.CreateSeal(seal)
	if err != nil {
		return nil, err
	}
	return seal, nil
}

func (sli *SealLI) GetAllTags() ([]*models.Tag, error) {
	return sli.sealService.GetAllTags()
}

func (sli *SealLI) AddTag(name string) (*models.Tag, error) {
	existing, err := sli.sealService.GetTagWithName(name)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("tag with name '%s' already exsits", name)
	}
	tag := &models.Tag{
		Name: name,
	}
	tag, err = sli.sealService.CreateTag(*tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
