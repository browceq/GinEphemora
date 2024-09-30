package main

import (
	"EphemoraApi/internal/controllers"
	"EphemoraApi/internal/middleware"
	"EphemoraApi/internal/repository"
	"EphemoraApi/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config %v. Shutting down...", err)
		return
	}

	dbDriver := config.Database.Driver
	dbUrl := "user=" + config.Database.User +
		" dbname=" + config.Database.DBName +
		" password=" + config.Database.Password +
		" host=" + config.Database.Host +
		" sslmode=" + config.Database.SSLMode

	server := gin.Default()

	mid := middleware.NewMiddleware()

	userRepo := repository.NewUserRepo(dbDriver, dbUrl)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, mid)

	leaderboardRepo := repository.NewLeaderboardRepo(dbDriver, dbUrl)
	leaderboardService := services.NewLeaderboardService(leaderboardRepo)
	leaderboardController := controllers.NewLeaderboardController(leaderboardService)

	server.POST("/users/signUp", userController.SignUp)
	server.POST("/users/login", userController.Login)

	auth := server.Group("/")
	auth.Use(mid.AuthMiddleware())
	{
		auth.PUT("/leaderboard/update", leaderboardController.UpdateRecord)
		auth.GET("/leaderboard/get", leaderboardController.GetLeaderboard)
	}

	err = server.Run(":" + config.Server.Port)
	if err != nil {
		return
	}

}
