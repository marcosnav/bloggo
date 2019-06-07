package router

import (
	ctrl "github.com/marcosnav/bloggo/controller"

	"github.com/gin-gonic/gin"
)

func setupControllers(router *gin.Engine) {
	router.GET("/", ctrl.Index)
}

// Setup all controllers and run the router
func Run(addr ...string) {
	router := gin.Default()
	setupControllers(router)
	router.Run(addr...)
}
