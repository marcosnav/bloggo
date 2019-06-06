package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash struct
type BcryptHash struct{}

// Generate bcrypt hash for given password
func (c BcryptHash) Generate(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare hashed password with plain password
// returns nil on success
func (c BcryptHash) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
