// Package models defines SEAaLS internal model representaions
// These are similar to the database models, but with better defined
// relations for seals and tags
package models

import (
	"time"
)

type Seal struct {
	ID        int64  `json:"id"`
	Path      string `json:"path"`
	Tags      []*Tag
	MimeType  string    `json:"mime_type"`
	CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TagStat struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
