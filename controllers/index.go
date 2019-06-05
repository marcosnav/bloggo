package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appservices "marcosn.com/services/application"
)

func Index(c *gin.Context) {
	token := appservices.NewJWTToken()

	// test validate token
	// valid := appservices.ValidateJWTToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAyOTYwODEsImlhdCI6MTU1OTY5MTI4MX0.mNUqSs93v7B74bZlUx_9U1LdRsnVj6lFN8vpNTvSU_s")
	// fmt.Println(valid)

	c.JSON(http.StatusOK, gin.H{
		"message": token,
	})
}
