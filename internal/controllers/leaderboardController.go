package controllers

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeaderboardConntroller interface {
	UpdateRecord(c *gin.Context)
	GetLeaderboard(c *gin.Context)
}

type leaderboardController struct {
	lrService services.LeaderboardService
}

func NewLeaderboardController(service services.LeaderboardService) LeaderboardConntroller {
	return &leaderboardController{service}
}

func (lrController *leaderboardController) UpdateRecord(c *gin.Context) {
	var newRecord models.RecordDTO
	if err := c.ShouldBindJSON(&newRecord); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	userEmail, ok := c.Get("user_email")
	if !ok {
		c.JSON(400, gin.H{"error": "No email in your token"})
		return
	}
	userEmailStr, ok := userEmail.(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid email in your token"})
		return
	}
	if newRecord.Email != userEmailStr {
		c.JSON(400, gin.H{"error": "Suspicious activity"})
		return
	}

	err := lrController.lrService.UpdateRecord(newRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update your record"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}

func (lrController *leaderboardController) GetLeaderboard(c *gin.Context) {
	leaderboard, err := lrController.lrService.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get leaderboard"})
		return
	}
	c.JSON(http.StatusOK, leaderboard)
}
