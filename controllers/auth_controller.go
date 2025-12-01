package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshitizgajurel/go_first_project/config"
	"github.com/kshitizgajurel/go_first_project/models"
	"github.com/kshitizgajurel/go_first_project/utils"
)


type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email"})
		return 
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email"})
		return
	}

	// just match text for now:
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
	}

	token, err := utils.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token creation failed"})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}