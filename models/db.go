package models

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitialiseDB(dbname string) (*sql.DB, error) {
	dbPath, err := filepath.Abs(dbname)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connecting to %s\n", dbPath)
	_, err = os.Stat(dbPath)
	if err != nil {
		// Create the database
		fmt.Printf("Database %s does not exist. Creating it...\n", dbPath)
		_, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
