package appservices

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

func (c *Hash) Generate(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *Hash) Compare(a string, b string) error {
	return bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
}
