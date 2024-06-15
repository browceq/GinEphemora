package main

import (
	"EphemoraApi/internal/controllers"
	"EphemoraApi/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/users/signUp", controllers.SignUp)
	server.POST("/users/login", controllers.Login)

	auth := server.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.PUT("/leaderboard/update", controllers.UpdateRecord)
		auth.GET("/leaderboard/get", controllers.GetLeaderboard)
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
