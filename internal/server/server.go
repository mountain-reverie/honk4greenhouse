package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mountain-reverie/honk4greenhouse/internal/handlers"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.Home)
	r.GET("/health", handlers.HealthCheck)

	return r
}

func Start(r *gin.Engine, port string) error {
	log.Printf("Starting server on port %s", port)
	return r.Run(":" + port)
}