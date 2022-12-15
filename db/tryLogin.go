package db

import (
	"github.com/lucassingh/go-backend/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin */
func TryLogin(email string, password string) (models.User, bool) {

	user, findUser, _ := CheckUserExist(email)

	if !findUser {
		return user, false
	}

	passworddBytes := []byte(password)

	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passworddBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
