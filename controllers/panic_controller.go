package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendPanic(c *gin.Context) {
	// Dummy handler
	c.JSON(http.StatusOK, gin.H{"message": "Panic button received!"})
}
