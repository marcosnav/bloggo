package main

import (
	"github.com/marcosnav/bloggo/config"
	"github.com/marcosnav/bloggo/router"
	"github.com/marcosnav/bloggo/tool"
)

// TODO: Move application to its own module/package and use DI to pass config and router
func main() {
	fileHandler := tool.LoadFileHandler()
	config.LoadWith(fileHandler)
	port, err := config.Get("APP_PORT")
	if err != nil {
		port = ":8080"
	}
	router.Run(port)
}
