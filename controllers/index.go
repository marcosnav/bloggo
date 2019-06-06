package controllers

import (
	"net/http"

	util "bloggo/util"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	jwt := util.JWT{Key: "secret"}
	token, _ := jwt.NewToken()

	c.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
