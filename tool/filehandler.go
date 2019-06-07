package tool

import (
	"fmt"
	"io/ioutil"
	"sync"

	errmsg "bloggo/error"
)

// Using a single instance of fileHandler (Singleton), as the only provider for file handling
// implementing sync.Once to prevent multi-thread race conditions
type fileHandler struct{}

type FileHandler interface {
	Read(file string) (string, error)
}

var once sync.Once
var instance fileHandler

// Get/Load the unique instance of the File Handler
func LoadFileHandler() fileHandler {
	// Thread safe, prevent race condition
	once.Do(func() {
		instance = fileHandler{}
	})
	return instance
}

//  Read a given file
func (fh fileHandler) Read(file string) (string, error) {
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf(errmsg.UnreadableFile, file)
	}
	return string(fileData), nil
}
