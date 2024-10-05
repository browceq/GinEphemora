package handler

import (
	"EphemoraApi/internal/models"
	mock_service "EphemoraApi/internal/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"

	"testing"
)

//	Unit-тесты userHndler со сгенерированными моками через gomock

func TestUserHandler_SignUpHandler(t *testing.T) {

	type mockBehaviour func(s *mock_service.MockUserService, user models.User)

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test","password":"qwerty","nickname":"nick"}`,
			inputUser: models.User{
				Email:    "test",
				Password: "qwerty",
				Nickname: "nick",
			},
			mockBehaviour: func(s *mock_service.MockUserService, user models.User) {
				s.EXPECT().AddUser(user).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User added successfully"}`,
		},
		{
			name:      "Invalid JSON",
			inputBody: `{"email":"test","nickname":"nick"}`, // пропущено поле nickname
			inputUser: models.User{
				Email:    "test",
				Nickname: "nick",
			},
			mockBehaviour: func(s *mock_service.MockUserService, user models.User) {
				// т.к. json невалидный, метод сервиса не вызывается
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"Invalid JSON"}`,
		},
		{
			name:      "Internal Server Error",
			inputBody: `{"email":"test","password":"qwerty","nickname":"nick"}`,
			inputUser: models.User{
				Email:    "test",
				Password: "qwerty",
				Nickname: "nick",
			},
			mockBehaviour: func(s *mock_service.MockUserService, user models.User) {
				s.EXPECT().AddUser(user).Return(errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error:": "Failed to add user. Maybe email is already taken"}`,
		},
	}

	//	запускаем саб-тесты для каждого testCase
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			//	Создаем контроллер для мока
			c := gomock.NewController(t)
			defer c.Finish()

			//	Создаем мок сервиса
			service := mock_service.NewMockUserService(c)

			//	Настраиваем поведение мока
			testCase.mockBehaviour(service, testCase.inputUser)

			//	Создаем тестовый контекст gin
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/signup", strings.NewReader(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			// Инициализируем хендлер
			userHandler := &userHandler{userService: service}

			// Вызываем тестируемый метод
			userHandler.SignUp(ctx)

			// Проверяем статус код
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

			// Проверяем тело ответа
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}
