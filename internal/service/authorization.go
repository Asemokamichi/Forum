package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Asemokamichi/Forum/internal/model"
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

func (servise *Service) GetUser(user model.User) (model.User, error) {
	hashPass, err := genereteHashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}

	user.HashedPassword = hashPass

	checkUser, err := servise.repository.GetUser(user.Username)
	if err != nil {
		return model.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.HashedPassword), []byte(user.Password)); err != nil {
		return model.User{}, fmt.Errorf("NOT HashedPassword")
	}

	return checkUser, nil
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

