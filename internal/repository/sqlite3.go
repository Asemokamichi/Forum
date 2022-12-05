package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDataBase(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/"+dbName)
	if err != nil {
		log.Fatal("Create Open error: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Create Open error: ", err)
		return nil, err
	}

	return db, nil
}

func (db *Repository) Create() error {
	query := `
		CREATE TABLE USER(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE,
			Email TEXT NOT NULL UNIQUE,
			Password TEXT NOT NULL
		);
	`

	db.db.Prepare(query)
	return nil
}
