package main

import (
	"EphemoraApi/internal/handler"
	"EphemoraApi/internal/middleware"
	"EphemoraApi/internal/repository"
	"EphemoraApi/internal/service"
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
	userService := service.NewUserService(userRepo)
	userController := handler.NewUserHandler(userService, mid)

	leaderboardRepo := repository.NewLeaderboardRepo(dbDriver, dbUrl)
	leaderboardService := service.NewLeaderboardService(leaderboardRepo)
	leaderboardController := handler.NewLeaderboardHandler(leaderboardService)

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
