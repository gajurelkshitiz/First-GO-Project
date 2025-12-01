package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/config"
	"github.com/kshitizgajurel/go_first_project/models"
)

// type Product struct {
// 	gorm.Model
// 	Code string
// 	Price uint
// }

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

	// HomePage for minimalistic app
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to minimalistic app.",
		})
	})


	database.connectDB()

	// Auto migrate tables
	database.DB.AutoMigrate(&models.User{})




	// server starts on localhost:8080 by default
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}