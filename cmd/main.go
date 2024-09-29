package main

import (
	"EphemoraApi/internal/controllers"
	"EphemoraApi/internal/middleware"
	"EphemoraApi/internal/repository"
	"EphemoraApi/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	userRepo := repository.NewUserRepo()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	leaderboardRepo := repository.NewLeaderboardRepo()
	leaderboardService := services.NewLeaderboardService(leaderboardRepo)
	leaderboardController := controllers.NewLeaderboardController(leaderboardService)

	mid := middleware.NewMiddleware()

	server.POST("/users/signUp", userController.SignUp)
	server.POST("/users/login", userController.Login)

	auth := server.Group("/")
	auth.Use(mid.AuthMiddleware())
	{
		auth.PUT("/leaderboard/update", leaderboardController.UpdateRecord)
		auth.GET("/leaderboard/get", leaderboardController.GetLeaderboard)
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
