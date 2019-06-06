package main

import (
	"bloggo/config"
	"bloggo/router"
	"bloggo/util"
)

// TODO: Move application to its own module/package and use DI to pass config and router
func main() {
	fileHandler := util.LoadFileHandler()
	config.LoadWith(fileHandler)
	port, err := config.Get("APP_PORT")
	if err != nil {
		port = ":8080"
	}
	router.Run(port)
}
