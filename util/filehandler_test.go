package util

import (
	"testing"

	I "bloggo/interfaces"

	"github.com/stretchr/testify/assert"
)

func TestLoadFileHandler(t *testing.T) {
	fh := LoadFileHandler()
	assert.Implements(t, (*I.FileHandler)(nil), fh, "should implement FileHandler")
}

func TestRead(t *testing.T) {
	fh := LoadFileHandler()
	content, _ := fh.Read("../fixtures/file.txt")
	assert.Equal(t, "Demo file for testing purposes", content, "should read files just fine")
}

func TestReadError(t *testing.T) {
	fh := LoadFileHandler()
	expectedErr := "Unable to read file unexistent_file.txt"
	_, err := fh.Read("unexistent_file.txt")
	assert.EqualError(t, err, expectedErr, "should return the error")
}
