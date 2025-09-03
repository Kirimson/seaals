package service

import (
	"context"
	"database/sql"
	"seaals/db"
	"seaals/models"
)

// SealService will interact with the underlying database to read Seals
type SealService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewSealService(Db *sql.DB) *SealService {
	q := db.New(Db)
	return &SealService{db: Db, queries: q}
}

func (ss *SealService) GetAllSeals() ([]models.Seal, error) {
	return nil, nil
}

func (ss *SealService) GetSealByID(id int64) (*models.Seal, error) {
	ctx := context.Background()
	seal, err := convertSeal(ss.queries.GetSeal(ctx, id))
	if err != nil {
		return nil, err
	}
	return seal, nil
}

func (ss *SealService) GetRandomSeal() (*models.Seal, error) {
	ctx := context.Background()
	// Get a random seal from the DB and
	// convert it to a 'regular' seaals seal
	seal, err := convertSeal(ss.queries.RandomSeal(ctx))
	if err != nil {
		return nil, err
	}

	// Get tags for this seal
	tags, err := ss.queries.ListSealTags(ctx, seal.ID)
	if err != nil {
		return nil, err
	}

	for _, t := range tags {
		seal.Tags = append(seal.Tags, models.Tag{Name: t.Name})
	}
	return seal, err
}

func (ss *SealService) GetRandomSealWithTag(tag string) (*models.Seal, error) {
	ctx := context.Background()
	// Get a random seal from the DB and
	// convert it to a 'regular' seaals seal
	seal, err := convertSeal(ss.queries.RandomSealWithTag(ctx, tag))
	if err != nil {
		return nil, err
	}

	// Get tags for this seal
	tags, err := ss.queries.ListSealTags(ctx, seal.ID)
	if err != nil {
		return nil, err
	}

	for _, t := range tags {
		seal.Tags = append(seal.Tags, models.Tag{Name: t.Name})
	}
	return seal, err
}

func (ss *SealService) CreateSeal(seal *models.Seal) (*models.Seal, error) {
	ctx := context.Background()
	args := db.CreateSealParams{
		Path:     seal.Path,
		MimeType: seal.MimeType,
	}
	newSeal, err := convertSeal(ss.queries.CreateSeal(ctx, args))
	if err != nil {
		return nil, err
	}
	return newSeal, nil
}
