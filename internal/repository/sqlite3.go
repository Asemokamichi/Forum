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
		CREATE TABLE IF NOT EXISTS USER(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE,
			Email TEXT NOT NULL UNIQUE,
			Password TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS SESSION(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			UserID INTEGER NOT NULL UNIQUE,
			UUID TEXT NOT NULL,
			ExpDate DATATIME NOT NULL
		);
	`
	if _, err := db.db.Exec(query); err != nil {
		return err
	}
	return nil
}
