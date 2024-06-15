package controllers

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/services"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.AddUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "Failed to add user. Maybe email is already taken"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func Login(c *gin.Context) {

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))
	var user models.UserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := services.Login(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login. Check your email and password"})
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &jwt.MapClaims{
		"user_email": user.Email,
		"exp":        expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successful login.", "token": tokenString})

}
