package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// create a gin router with default middleware
	r := gin.Default()

	// Define a simple GET Endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK,gin.H{
			"message" : "pong",
		})
	})

	// server starts on localhost:8080 by default
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}