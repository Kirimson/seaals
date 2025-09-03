package service

import (
	"seaals/models"
)

// GetTagByName returns a tag from a given name. If the Tag does not exist
// nil will be returned, and Gorm will not print an error to stderr
func (ss *SealService) GetTagByName(name string) (*models.Tag, error) {
	return nil, nil
}

func (ss *SealService) CreateTag(tag models.Tag) (*models.Tag, error) {
	return nil, nil
}

func (ss *SealService) GetAllTags() ([]models.Tag, error) {
	return nil, nil
}
