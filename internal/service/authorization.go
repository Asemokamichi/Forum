package service

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Asemokamichi/Forum/internal/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (servise *Service) CreateUser(user model.User) error {
	if !CheckEmail(user.Email) || !CheckPassword(user.Password) {
		return fmt.Errorf("Incorrect user info")
	}

	password, err := genereteHashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password
	return servise.repository.CreateUser(user)

}

func (servise *Service) CreateSession(username string, password string) (model.Session, error) {
	user, err := servise.repository.GetUser(username)
	if err != nil {
		return model.Session{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return model.Session{}, fmt.Errorf("NOT HashedPassword")
	}

	var session model.Session = model.Session{
		UserID:  user.ID,
		UUID:    uuid.NewString(),
		ExpDate: time.Now().Add(6 * time.Hour),
	}

	if err := servise.repository.CreateSession(session); err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (service *Service) GetUserSession(token string) (model.User, error) {
	session, err := service.repository.GetSession(token)
	if err != nil {
		return model.User{}, err
	}

	user, err := service.repository.GetUserID(session.UserID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func CheckPassword(password string) bool {
	numbers := "0123456789"
	lowerCase := "qwertyuiopasdfghjklzxcvbnm"
	upperCase := "QWERTYUIOPASDFGHJKLZXCVBNM"
	symbols := "@#$%^&*()_-+={[}]|\\:;<,>.?/"

	if !checkContains(password, numbers) || !checkContains(password, lowerCase) || !checkContains(password, upperCase) || !checkContains(password, symbols) {
		return false
	}

	for _, w := range symbols {
		if w < 32 || w > 126 {
			return false
		}
	}
	return true
}

func CheckEmail(email string) bool {
	return regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email)
}

func checkContains(s, checkSymbols string) bool {
	for _, w := range checkSymbols {
		if strings.Contains(s, string(w)) {
			return true
		}
	}
	return false
}

func genereteHashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(pass), err
}
