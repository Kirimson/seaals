package service

import (
	"context"
	"seaals/models"

	"gorm.io/gorm"
)

// GetTagByName returns a tag from a given name. If the Tag does not exist
// nil will be returned, and Gorm will not print an error to stderr
func (ss *SealService) GetTagByName(name string) (*models.Tag, error) {
	ctx := context.Background()
	result := gorm.WithResult()
	tag, err := gorm.G[models.Tag](ss.db, result).Where("name = ?", name).Limit(1).Find(ctx)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected > 0 {
		return &tag[0], nil
	}
	return nil, nil
}

func (ss *SealService) CreateTag(tag models.Tag) (*models.Tag, error) {
	ctx := context.Background()
	result := gorm.WithResult()
	err := gorm.G[models.Tag](ss.db, result).Create(ctx, &tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}
