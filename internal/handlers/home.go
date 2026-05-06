package handlers

import (
	_ "embed"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed templates/home.html
var homeHTML string

func Home(c *gin.Context) {
	acceptHeader := c.GetHeader("Accept")
	
	if strings.Contains(acceptHeader, "text/html") || acceptHeader == "" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, homeHTML)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! Welcome to honk4greenhouse - your greenhouse design companion",
			"version": "1.0.0",
		})
	}
}