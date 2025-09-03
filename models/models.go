// Package models defines SEAaLS internal model representaions
// These are similar to the database models, but with better defined
// relations for seals and tags
package models

type Seal struct {
	ID       int64  `json:"id"`
	Path     string `json:"path"`
	Tags     []Tag
	MimeType string `json:"mime_type"`
}

type Tag struct {
	Name string `json:"name"`
}
