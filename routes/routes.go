package routes

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/controllers"
	"github.com/kshitizgajurel/go_first_project/middleware"
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
		api.POST("/login", controllers.Login)
		api.POST("/user", controllers.CreateUser)
		api.GET("/users", middleware.AuthMiddleware(),  controllers.GetUsers)
	}
}