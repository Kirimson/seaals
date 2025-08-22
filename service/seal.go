package service

import (
	"context"
	"math/rand"
	"seaals/models"

	"gorm.io/gorm"
)

// SealService will interact with the underlying database to read Seals
type SealService struct {
	db *gorm.DB
}

func NewSealService(Db *gorm.DB) *SealService {
	return &SealService{db: Db}
}

func (ss *SealService) GetRandomSeal() (*models.Seal, error) {
	ctx := context.Background()
	total, err := gorm.G[models.Seal](ss.db).Count(ctx, "path")
	if err != nil {
		return nil, err
	}
	randOffset := rand.Intn(int(total))

	seal, err := gorm.G[models.Seal](ss.db).Offset(randOffset).First(ctx)
	if err != nil {
		return nil, err
	}

	return &seal, nil
}

func (ss *SealService) GetSealByID(id uint) (*models.Seal, error) {
	ctx := context.Background()
	seal, err := gorm.G[models.Seal](ss.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &seal, nil
}

func (ss *SealService) GetRandomSealByTag(tag string) (*models.Seal, error) {
	ctx := context.Background()
	tagTotal, err := gorm.G[models.Seal](ss.db).Where("tag = ?", tag).Count(ctx, "id")
	if err != nil {
		return nil, err
	}

	randOffset := rand.Intn(int(tagTotal))
	seal, err := gorm.G[models.Seal](ss.db).Where("tag = ?", tag).Offset(randOffset).First(ctx)
	if err != nil {
		return nil, err
	}
	return &seal, nil
}

func (ss *SealService) CreateSeal(seal *models.Seal) (*models.Seal, error) {
	ctx := context.Background()
	result := gorm.WithResult()
	err := gorm.G[models.Seal](ss.db, result).Create(ctx, seal)
	if err != nil {
		return nil, err
	}
	return seal, nil
}

func (ss *SealService) GetAllSeals() ([]models.Seal, error) {
	ctx := context.Background()
	seals, err := gorm.G[models.Seal](ss.db).Preload("Tags", nil).Find(ctx)
	if err != nil {
		return nil, err
	}
	return seals, nil
}
