// Package models defines SEAaLS database models
// Different client implementations will interact with
// these models
package models

import "gorm.io/gorm"

type Seal struct {
	gorm.Model
	Path string
	Tags []Tag `gorm:"many2many:seal_tags;"`
}
