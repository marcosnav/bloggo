package util

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Key string
}

// Generate new token with key(secret)
func (j JWT) NewToken() (string, error) {
	now := time.Now()
	exp := now.AddDate(0, 0, 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": now.Unix(),
		"exp": exp.Unix(),
	})
	tokenString, err := token.SignedString([]byte(j.Key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Validate token
func (j JWT) ValidateToken(token string) bool {
	parsedToken, err := jwt.Parse(token, j.keyFuncOnParse)
	if err != nil {
		return false
	}
	return parsedToken.Valid
}

func (j JWT) keyFuncOnParse(tkn *jwt.Token) (interface{}, error) {
	if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
	}
	return []byte(j.Key), nil
}
