package repository

import (
	"github.com/Asemokamichi/Forum/internal/model"
)

func (db *Repository) CreateUser(user model.User) error {
	query := `
		INSERT INTO USER (Username, Email, Password) VALUES ($1, $2, $3);
	`

	if _, err := db.db.Exec(query, user.Username, user.Email, user.Password); err != nil {
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
		return model.User{}, err
	}
	return user, nil
}

func (db *Repository) GetUserID(ID int) (model.User, error) {
	query := `
		SELECT ID, Username, Email, Password FROM USER WHERE ID = ?;
	`
	var user model.User
	if err := db.db.QueryRow(query, ID).Scan(&user.ID, &user.Username, &user.Email, &user.HashedPassword); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (db *Repository) CreateSession(user model.Session) error {
	query := `
		INSERT INTO SESSION (UserID, UUID, ExpDate) VALUES ($1, $2, $3);
	`

	if _, err := db.db.Exec(query, user.UserID, user.UUID, user.ExpDate); err != nil {
		return err
	}

	return nil
}

func (db *Repository) GetSession(UUID string) (model.Session, error) {
	query := `
		SELECT ID, UserID, UUID FROM SESSION WHERE UUID = ?;
	`
	var session model.Session
	if err := db.db.QueryRow(query, UUID).Scan(&session.ID, &session.UserID, &session.UUID); err != nil {
		return model.Session{}, err
	}
	return session, nil
}
