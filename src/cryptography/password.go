package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, dbPassword string) error {

	passwordInString := []byte(password)
	dbPasswordInString := []byte(dbPassword)

	return bcrypt.CompareHashAndPassword(passwordInString, dbPasswordInString)
}
