package appservices

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"marcosn.com/config"
)

func signingKey() []byte {
	secret := config.Get("APP_SECRET")
	return []byte(secret)
}

func validateParsedToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return signingKey(), nil
}

func NewJWTToken() string {
	now := time.Now()
	exp := now.AddDate(0, 0, 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": now.Unix(),
		"exp": exp.Unix(),
	})
	tokenString, err := token.SignedString(signingKey())
	if err != nil {
		// Handle error
	}
	return tokenString
}

func ValidateJWTToken(token string) bool {
	parsedToken, err := jwt.Parse(token, validateParsedToken)
	if err != nil {
		return false
	}
	return parsedToken.Valid
}
