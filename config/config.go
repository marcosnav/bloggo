package config

import (
	"io/ioutil"
	"reflect"

	"github.com/BurntSushi/toml"
)

type Config struct {
	APP_PORT   string
	APP_SECRET string
}

var config Config

func init() {
	configFile, err := ioutil.ReadFile("config.toml")
	if err != nil {
		// Handle error
	}
	configData := string(configFile)
	if _, err := toml.Decode(configData, &config); err != nil {
		// handle error
	}
}

func Get(field string) string {
	r := reflect.ValueOf(config)
	f := reflect.Indirect(r).FieldByName(field)
	if !f.IsValid() {
		// Handle invalid/unexistent config field
	}
	return f.String()
}
