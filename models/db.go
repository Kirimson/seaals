package models

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitialiseDB(dbname string) (*sql.DB, error) {
	dbPath, err := filepath.Abs(dbname)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connecting to %s\n", dbPath)
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
