package controllers

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "Failed to add user. Maybe email is already taken"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func Login(c *gin.Context) {
	var user models.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := services.Login(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login. Check your email and password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successful login"})

}
