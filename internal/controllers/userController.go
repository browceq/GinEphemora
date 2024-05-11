package controllers

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new user
// @Description Add a new user to the system
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User object to be added"
// @Success 200 {object} gin.H{"message": "User added successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid JSON"}
// @Failure 500 {object} gin.H{"error": "Failed to add user"}
// @Router /users [post]
func AddUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	err := services.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "Failed to add user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}
