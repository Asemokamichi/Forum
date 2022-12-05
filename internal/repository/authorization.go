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

func (db *Repository) GetUser(username string) (model.User, error) {
	query := `
		SELECT ID, Username, Email, Password FROM USER WHERE Username = ?;
	`
	var user model.User
	if err := db.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword); err != nil {
		log.Fatal(err)
	}
	return user, nil
}
