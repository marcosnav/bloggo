package util

import (
	"testing"

	I "bloggo/interfaces"

	"github.com/stretchr/testify/assert"
)

var hash string

func TestNewHasher(t *testing.T) {
	hasher := BcryptHash{}
	assert.Implements(t, (*I.BcryptHasher)(nil), hasher, "should implement BcryptHasher")
}

func TestGenerate(t *testing.T) {
	hasher := BcryptHash{}
	h, err := hasher.Generate("a&password!")
	hash = h
	assert.Nil(t, err, "without errors")
	assert.IsType(t, (string)(""), h, "returns a hash as string")
	assert.NotEmpty(t, h, "must not be empty")
}

func TestCompare(t *testing.T) {
	hasher := BcryptHash{}
	err := hasher.Compare(hash, "a&password!")
	assert.Nil(t, err, "without errors")
}

func TestCompareError(t *testing.T) {
	hasher := BcryptHash{}
	err := hasher.Compare(hash, "another&password")
	assert.Error(t, err, "returns hash mismatch error")
}
