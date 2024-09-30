package controllers

import (
	"EphemoraApi/internal/middleware"
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService services.UserService
	middleware  middleware.Middleware
}

func NewUserController(userService services.UserService, middleware middleware.Middleware) UserController {
	return &userController{userService, middleware}
}

func (uC *userController) SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	err := uC.userService.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "Failed to add user. Maybe email is already taken"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func (uC *userController) Login(c *gin.Context) {

	var user models.UserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := uC.userService.Login(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login. Check your email and password"})
		return
	}

	tokenString, err := uC.middleware.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "Successful login."})
}
