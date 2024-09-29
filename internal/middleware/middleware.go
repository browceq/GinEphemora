package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthMiddleware() gin.HandlerFunc
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
