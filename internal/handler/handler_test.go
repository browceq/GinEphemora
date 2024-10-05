package handler

import (
	mock_middleware "EphemoraApi/internal/middleware/mocks"
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

			//	Настраиваем поведение мока:
			//	Вызываем ф-ию mockBehaviour с аргументами service и user
			//	Функция задает поведение конкретному тесту
			testCase.mockBehaviour(service, testCase.inputUser)

			//	Создаем тестовый контекст gin
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/signup", strings.NewReader(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			// Инициализируем хендлер
			userHandler := &userHandler{userService: service}

			// Вызываем тестируемый метод, передаем тестовый контекст
			userHandler.SignUp(ctx)

			// Проверяем статус код
			assert.Equal(t, testCase.expectedStatusCode, w.Code)

			// Проверяем тело ответа
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}

func TestUserHandler_Login(t *testing.T) {
	type ServiceMockBehaviour func(s *mock_service.MockUserService, user models.UserDTO)
	type MiddlewareMockBehaviour func(m *mock_middleware.MockMiddleware, user models.UserDTO)

	testTable := []struct {
		name                    string
		inputBody               string
		inputUser               models.UserDTO
		ServiceMockBehaviour    ServiceMockBehaviour
		MiddlewareMockBehaviour MiddlewareMockBehaviour
		expectedStatusCode      int
		expectedResponseBody    string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test","password":"qwerty"}`,
			inputUser: models.UserDTO{
				Email:    "test",
				Password: "qwerty",
			},
			ServiceMockBehaviour: func(s *mock_service.MockUserService, user models.UserDTO) {
				s.EXPECT().Login(user).Return(nil)
			},
			MiddlewareMockBehaviour: func(m *mock_middleware.MockMiddleware, user models.UserDTO) {
				m.EXPECT().GenerateToken(user).Return("token", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message": "Successful login."}`,
		},
		{
			name:      "Invalid JSON",
			inputBody: `{"email":"test"}`,
			inputUser: models.UserDTO{
				Email: "test",
			},

			ServiceMockBehaviour: func(s *mock_service.MockUserService, user models.UserDTO) {

			},
			MiddlewareMockBehaviour: func(m *mock_middleware.MockMiddleware, user models.UserDTO) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error": "Invalid JSON"}`,
		},
		{
			name:      "Failed to generate token",
			inputBody: `{"email":"test","password":"qwerty"}`,
			inputUser: models.UserDTO{
				Email:    "test",
				Password: "qwerty",
			},
			ServiceMockBehaviour: func(s *mock_service.MockUserService, user models.UserDTO) {
				s.EXPECT().Login(user).Return(nil)
			},
			MiddlewareMockBehaviour: func(m *mock_middleware.MockMiddleware, user models.UserDTO) {
				m.EXPECT().GenerateToken(user).Return("token", errors.New("middleware error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error": "Failed to generate token"}`,
		},
		{
			name:      "Failed to login",
			inputBody: `{"email":"test","password":"qwerty"}`,
			inputUser: models.UserDTO{
				Email:    "test",
				Password: "qwerty",
			},
			ServiceMockBehaviour: func(s *mock_service.MockUserService, user models.UserDTO) {
				s.EXPECT().Login(user).Return(errors.New("service error"))
			},
			MiddlewareMockBehaviour: func(m *mock_middleware.MockMiddleware, user models.UserDTO) {
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error": "Failed to login. Check your email and password"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockUserService(c)
			middleware := mock_middleware.NewMockMiddleware(c)

			testCase.ServiceMockBehaviour(service, testCase.inputUser)
			testCase.MiddlewareMockBehaviour(middleware, testCase.inputUser)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			req := httptest.NewRequest("POST", "/login", strings.NewReader(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req

			userHandler := &userHandler{userService: service, middleware: middleware}

			userHandler.Login(ctx)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestLeaderboarHandler_GetLeaderboard(t *testing.T) {

	type LeaderboardMockBehaviour func(s *mock_service.MockLeaderboardService)
	//type MiddlewareMockBehaviour func(m *mock_middleware.MockMiddleware)

	testTable := []struct {
		name                     string
		LeaderboardMockBehaviour LeaderboardMockBehaviour
		expectedStatusCode       int
		expectedResponseBody     string
		//MiddlewareMockBehaviour MiddlewareMockBehaviour
	}{
		{
			name: "OK",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService) {
				s.EXPECT().GetLeaderboard().Return([]models.LeaderboardEntry{
					{"player1", 10},
					{"player2", 9},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"nickname":"player1","record":10},{"nickname":"player2","record":9}]`,
		},
		{
			name: "Internal Server Error",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService) {
				s.EXPECT().GetLeaderboard().Return(nil, errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"Failed to get leaderboard"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_service.NewMockLeaderboardService(c)

			testCase.LeaderboardMockBehaviour(service)

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			lrHandler := &leaderboarHandler{lrService: service}

			lrHandler.GetLeaderboard(ctx)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}

func TestLeaderboarHandler_UpdateRecord(t *testing.T) {
	type LeaderboardMockBehaviour func(s *mock_service.MockLeaderboardService, record models.RecordDTO)
	testTable := []struct {
		name                     string
		inputBody                string
		inputRecord              models.RecordDTO
		ctxEmailKey              interface{}
		LeaderboardMockBehaviour LeaderboardMockBehaviour
		expectedStatusCode       int
		expectedResponseBody     string
	}{
		{
			name:      "OK",
			inputBody: `{"email":"test","record":1}`,
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: "test",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {
				s.EXPECT().UpdateRecord(record).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message": "Record updated successfully"}`,
		},
		{
			name:      "Invalid JSON",
			inputBody: `{"email":"test","record":"1""}`, //record string, вместо int
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: "test",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {

			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error": "Invalid JSON"}`,
		},
		{
			name:      "No email in your token",
			inputBody: `{"email":"test","record":1}`,
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: nil,
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error": "No email in your token"}`,
		},
		{
			name:      "Invalid email in your token",
			inputBody: `{"email":"test","record":1}`,
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: 123,
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {

			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error": "Invalid email in your token"}`,
		},
		{
			name:      "Suspicious activity",
			inputBody: `{"email":"test","record":1}`,
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: "fake",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {

			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error": "Suspicious activity"}`,
		},
		{
			name:      "Failed to update your record",
			inputBody: `{"email":"test","record":1}`,
			inputRecord: models.RecordDTO{
				Email:  "test",
				Record: 1,
			},
			ctxEmailKey: "test",
			LeaderboardMockBehaviour: func(s *mock_service.MockLeaderboardService, record models.RecordDTO) {
				s.EXPECT().UpdateRecord(record).Return(errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message": "Failed to update your record"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			if testCase.ctxEmailKey != nil {
				ctx.Set("user_email", testCase.ctxEmailKey)
			}

			req := httptest.NewRequest("PUT", "/leaderboard/update", strings.NewReader(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")

			ctx.Request = req

			service := mock_service.NewMockLeaderboardService(c)
			testCase.LeaderboardMockBehaviour(service, testCase.inputRecord)

			lrHandler := &leaderboarHandler{lrService: service}

			lrHandler.UpdateRecord(ctx)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())

		})
	}

}
