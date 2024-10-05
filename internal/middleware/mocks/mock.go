// Code generated by MockGen. DO NOT EDIT.
// Source: middleware.go

// Package mock_middleware is a generated GoMock package.
package mock_middleware

import (
	models "EphemoraApi/internal/models"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockMiddleware is a mock of Middleware interface.
type MockMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockMiddlewareMockRecorder
}

// MockMiddlewareMockRecorder is the mock recorder for MockMiddleware.
type MockMiddlewareMockRecorder struct {
	mock *MockMiddleware
}

// NewMockMiddleware creates a new mock instance.
func NewMockMiddleware(ctrl *gomock.Controller) *MockMiddleware {
	mock := &MockMiddleware{ctrl: ctrl}
	mock.recorder = &MockMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleware) EXPECT() *MockMiddlewareMockRecorder {
	return m.recorder
}

// AuthMiddleware mocks base method.
func (m *MockMiddleware) AuthMiddleware() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMiddleware")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// AuthMiddleware indicates an expected call of AuthMiddleware.
func (mr *MockMiddlewareMockRecorder) AuthMiddleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMiddleware", reflect.TypeOf((*MockMiddleware)(nil).AuthMiddleware))
}

// GenerateToken mocks base method.
func (m *MockMiddleware) GenerateToken(user models.UserDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockMiddlewareMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockMiddleware)(nil).GenerateToken), user)
}
