package handlers

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashingHandler(value string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(b), err
}
