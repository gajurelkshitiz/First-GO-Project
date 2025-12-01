package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/config"
	"github.com/kshitizgajurel/go_first_project/models"
	"github.com/kshitizgajurel/go_first_project/routes"
)


func main() {
	// create a gin router with default middleware
	r := gin.Default()

	// register app routes
	routes.SetupRoutes(r)

	config.ConnectDB()

	
	
	// Auto migrate tables
	config.DB.AutoMigrate(&models.User{})

	// server starts on localhost:8080 by default
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}