package util

import (
	"testing"

	I "bloggo/interfaces"

	"github.com/stretchr/testify/assert"
)

var token string

func TestNewJWT(t *testing.T) {
	jwt := JWT{Key: "secret"}
	assert.Implements(t, (*I.TokenProvider)(nil), jwt, "should implement TokenProvider")
}

func TestNewToken(t *testing.T) {
	jwt := JWT{Key: "secret"}
	tkn, err := jwt.NewToken()
	token = tkn
	assert.IsType(t, (string)(""), token, "should return token string")
	assert.NotEmpty(t, tkn, "should not be not empty")
	assert.Nil(t, err, "without errors")
}

func TestValidateToken(t *testing.T) {
	jwt := JWT{Key: "secret"}
	res := jwt.ValidateToken(token)
	assert.Equal(t, true, res, "should be true")
}

func TestValidateTokenFalse(t *testing.T) {
	jwt := JWT{Key: "another$secret"}
	res := jwt.ValidateToken(token)
	assert.Equal(t, false, res, "should be false")
}
