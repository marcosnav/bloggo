// Package providing the controllers that will act as the handlers for upcoming outside requests
package controller

import (
	"net/http"

	"github.com/marcosnav/bloggo/tool"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	jwt := tool.JWT{Key: "secret"}
	token, _ := jwt.NewToken()

	c.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
