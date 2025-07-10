package main

import (
	"log"
	"os"

	"github.com/mountain-reverie/honk4greenhouse/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := server.New()
	
	if err := server.Start(r, port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
