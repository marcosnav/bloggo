package router

import (
	"github.com/gin-gonic/gin"
	ctrl "marcosn.com/controllers"
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
