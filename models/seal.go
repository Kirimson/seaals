package models

import "gorm.io/gorm"

type Seal struct {
	gorm.Model
	Path string
	Tags []Tag `gorm:"many2many:seal_tags;"`
}
