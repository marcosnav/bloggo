package config

import (
	"errors"
	"testing"

	errmsg "github.com/marcosnav/bloggo/error"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fileReaderMock struct {
	mock.Mock
}

func (m *fileReaderMock) Read(file string) (string, error) {
	args := m.Called(file)
	return args.String(0), args.Error(1)
}

func TestLoadWith(t *testing.T) {
	reader := new(fileReaderMock)
	reader.On("Read", "config.toml").Return("APP_PORT = \":3000\"", nil)

	err := LoadWith(reader)
	assert.Equal(t, nil, err, "should load config file")
}

func TestGet(t *testing.T) {
	port, _ := Get("APP_PORT")
	assert.Equal(t, ":3000", port, "should get config values")
}

func TestUnexistentGet(t *testing.T) {
	field, err := Get("UNEXISTENT_FIELD")
	expectErr := "The requested configuration field UNEXISTENT_FIELD does not exist on loaded config.toml"
	assert.Equal(t, "", field, "should get empty field")
	assert.EqualError(t, err, expectErr, "error must be returned")
}

func TestBadFormattedToml(t *testing.T) {
	reader := new(fileReaderMock)
	reader.On("Read", "config.toml").Return("= APP_PORT : \":3000\"", nil)

	err := readAndDecodeFile(reader)
	assert.EqualError(t, err, errmsg.UnparsableConfigFile, "error must be returned")
}

func TestConfigReadingError(t *testing.T) {
	reader := new(fileReaderMock)
	reader.On("Read", "config.toml").Return("", errors.New("Reader error"))

	err := readAndDecodeFile(reader)
	assert.EqualError(t, err, "Reader error", "error must be returned")
}
