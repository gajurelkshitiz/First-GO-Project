package routes

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// test routes
	router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Welcome to minimalistic app."})
    })


	api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/user", controllers.CreateUser)
	}
}