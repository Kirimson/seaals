package service

import (
	"context"
	"database/sql"
	"fmt"
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

func (ss *SealService) GetAllSeals() ([]*models.Seal, error) {
	ctx := context.Background()
	seals, err := convertSeals(ss.queries.ListSeals(ctx))
	if err != nil {
		return nil, err
	}

	return seals, nil
}

func (ss *SealService) GetSealByID(id int64) (*models.Seal, error) {
	ctx := context.Background()
	seal, err := convertSeal(ss.queries.GetSeal(ctx, id))
	if err != nil {
		return nil, err
	}
	if seal == nil {
		return seal, fmt.Errorf("no seal found with ID '%d'", id)
	}
	return seal, nil
}

func (ss *SealService) GetSealByPath(path string) (*models.Seal, error) {
	ctx := context.Background()
	// This is a **bit** hacky. Paths always end in a file extension, so use the provided path
	// and append .% to the LIKE query for an exact match disregarind the extension
	seal, err := convertSeal(ss.queries.GetSealByPath(ctx, fmt.Sprintf("%s.%%", path)))
	if err != nil {
		return nil, err
	}
	if seal == nil {
		return seal, fmt.Errorf("no seal found with ID '%s'", path)
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
		seal.Tags = append(seal.Tags, &models.Tag{Name: t.Name})
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
	if seal == nil {
		return nil, fmt.Errorf("no seal found with tag '%s'", tag)
	}

	// Get tags for this seal
	tags, err := ss.queries.ListSealTags(ctx, seal.ID)
	if err != nil {
		return nil, err
	}

	for _, t := range tags {
		seal.Tags = append(seal.Tags, &models.Tag{Name: t.Name})
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

	// Add tag associations for this seal
	for _, t := range seal.Tags {
		_, err := ss.queries.AddSealTag(ctx, db.AddSealTagParams{SealID: newSeal.ID, TagID: t.ID})
		if err != nil {
			return nil, err
		}
	}
	return newSeal, nil
}

func (ss *SealService) GetSealTags(seal *models.Seal) ([]*models.Tag, error) {
	ctx := context.Background()
	return convertTags(ss.queries.ListSealTags(ctx, seal.ID))
}

func (ss *SealService) DeleteSeal(id int64) error {
	ctx := context.Background()
	// Delete the seal
	err := ss.queries.DeleteSeal(ctx, id)
	if err != nil {
		return err
	}
	// Delete all tag associations with the seal to delete
	return ss.queries.DeleteSealTags(ctx, id)
}

func (ss *SealService) CountSeals() (int64, error) {
	ctx := context.Background()
	return ss.queries.CountSeals(ctx)
}
