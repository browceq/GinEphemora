package middleware

import (
	"EphemoraApi/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:generate mockgen -source=middleware.go -destination=mocks/mock.go

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
	GenerateToken(user models.UserDTO) (string, error)
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func (middleware *middleware) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Удаление префикса "Bearer " из заголовка
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Парсинг и проверка токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверка метода подписи токена
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Проверка и извлечение утверждений (claims) из токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_email", claims["user_email"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Передача следующей функции
		c.Next()
	}
}

func (middleware *middleware) GenerateToken(user models.UserDTO) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &jwt.MapClaims{
		"user_email": user.Email,
		"exp":        expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
