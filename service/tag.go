package service

import (
	"context"
	"seaals/models"
)

// GetTagWithName returns a tag from a given name. If the Tag does not exist
// nil will be returned, and Gorm will not print an error to stderr
func (ss *SealService) GetTagWithName(name string) (*models.Tag, error) {
	ctx := context.Background()
	return convertTag(ss.queries.GetTagWithName(ctx, name))
}

func (ss *SealService) CreateTag(tag models.Tag) (*models.Tag, error) {
	ctx := context.Background()
	newTag, err := convertTag(ss.queries.CreateTag(ctx, tag.Name))
	if err != nil {
		return nil, err
	}
	return newTag, nil
}

func (ss *SealService) GetAllTags() ([]*models.Tag, error) {
	ctx := context.Background()
	return convertTags(ss.queries.ListTags(ctx))
}
