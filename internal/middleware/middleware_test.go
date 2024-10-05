package middleware

import (
	"EphemoraApi/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMiddleware_GenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "secret")
	defer os.Unsetenv("JWT_SECRET")

	mw := NewMiddleware()

	user := models.UserDTO{Email: "test"}

	token, err := mw.GenerateToken(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestMiddleware_AuthMiddleware(t *testing.T) {

	os.Setenv("JWT_SECRET", "secret")
	defer os.Unsetenv("JWT_SECRET")

	mw := NewMiddleware()

	router := gin.Default()
	router.Use(mw.AuthMiddleware())

	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Protected"})
	})

	testTable := []struct {
		name               string
		token              string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "Missing Authorization Header",
			token:              "",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Authorization header required"}`,
		},
		{
			name:               "Invalid Token",
			token:              "Bearer invalidtoken",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Invalid token"}`,
		},
		{
			name:               "Valid Token",
			token:              "Bearer " + generateValidToken("test@example.com"),
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"message":"Protected"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set("Authorization", testCase.token)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponse, w.Body.String())

		})
	}

}

func generateValidToken(email string) string {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	claims := &jwt.MapClaims{
		"user_email": email,
		"exp":        jwt.NewNumericDate(time.Now().Add(15 * time.Minute)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}
