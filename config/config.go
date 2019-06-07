// Package that provides the application configuration in a Singleton manner
package config

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/BurntSushi/toml"

	errmsg "bloggo/error"
	"bloggo/tool"
)

const configFileName string = "config.toml"

// Using config as a Singleton, single source of truth for config fields
// implementing sync.Once to prevent multi-thread race conditions
type config struct {
	APP_PORT   string
	APP_SECRET string
}

var once sync.Once
var instance config

func readAndDecodeFile(fh tool.FileHandler) error {
	configData, readError := fh.Read(configFileName)
	if readError != nil {
		return readError
	}
	if _, err := toml.Decode(configData, &instance); err != nil {
		return errors.New(errmsg.UnparsableConfigFile)
	}
	return nil
}

// Load configuration from config.toml
func LoadWith(fh tool.FileHandler) error {
	var loadError error = nil
	// Thread safe, prevent race condition
	once.Do(func() {
		err := readAndDecodeFile(fh)
		if err != nil {
			loadError = err
		}
	})
	return loadError
}

// Get config field
func Get(field string) (string, error) {
	r := reflect.ValueOf(instance)
	f := reflect.Indirect(r).FieldByName(field)
	if !f.IsValid() {
		return "", fmt.Errorf(errmsg.UnexistentConfigField, field)
	}
	return f.String(), nil
}
