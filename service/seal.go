package service

import "gorm.io/gorm"

// SealService will interact with the underlying database to read Seals
type SealService struct {
	Db *gorm.DB
}

func NewSealService(Db *gorm.DB) *SealService {
	return &SealService{Db: Db}
}
