package controllers

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRecord(c *gin.Context) {
	var newRecord models.RecordDTO
	if err := c.ShouldBindJSON(&newRecord); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	userEmail, ok := c.Get("user_email")
	if !ok {
		c.JSON(400, gin.H{"error": "No email in your token"})
	}
	userEmailStr, ok := userEmail.(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid email in your token"})
	}
	if newRecord.Email != userEmailStr {
		c.JSON(400, gin.H{"error": "Suspicious activity"})
	}

	err := services.UpdateRecord(newRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update your record"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}

func GetLeaderboard(c *gin.Context) {
	leaderboard, err := services.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get leaderboard"})
	}
	c.JSON(http.StatusOK, leaderboard)
}
