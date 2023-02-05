package cryptography

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(password, dbPassword []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(dbPassword))
}
