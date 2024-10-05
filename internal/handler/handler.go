package handler

import "github.com/gin-gonic/gin"

type LeaderboardHandler interface {
	UpdateRecord(c *gin.Context)
	GetLeaderboard(c *gin.Context)
}

type UserHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}
