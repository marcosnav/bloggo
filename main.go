package main

import (
	"marcosn.com/config"
	"marcosn.com/router"
)

func main() {
	port := config.Get("APP_PORT")
	router.Run(port)
}
