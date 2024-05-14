package main

import (
	"EphemoraApi/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/users/addUser", controllers.AddUser)
	server.POST("/users/login", controllers.Login)

	server.PUT("/leaderboard/update", controllers.UpdateRecord)
	server.GET("/leaderboard/get", controllers.GetLeaderboard)

	server.Run(":8080")
}
