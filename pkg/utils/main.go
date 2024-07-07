package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(hashedPasswords, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswords), []byte(password))
	if err == nil {
		return true
	}
	return false

}
