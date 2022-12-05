package repository

import (
	"log"

	"github.com/Asemokamichi/Forum/internal/model"
)

func (db *Repository) CreateUser(user model.User) error {
	query := `
		INSERT INTO USER (Username, Email, Password) VALUES ($1, $2, $3);
	`

	if _, err := db.db.Exec(query, user.Username, user.Email, user.Password); err != nil {
		log.Fatal("Error Insert User: ", err)
		return err
	}

	return nil
}

func (db *Repository) GetUser(user model.User) error {
	query := `
		SELECT COUNT(*) FROM USER WHERE Username = ?;
	`
	var count int
	if err := db.db.QueryRow(query, user.Username, user.Password).Scan(&count); count!=0{
		log.Fatal(err)
	}

	if count 
}
