package main

import (
	"EphemoraApi/internal/controllers"

	_ "EphemoraApi/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ephemora API
// @version 1.1
// @host localhost:8080
func main() {
	server := gin.Default()

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.POST("/users", controllers.AddUser)

	server.Run(":8080")
}
