package models

import "database/sql"

func InitialiseDB(dbname string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}
