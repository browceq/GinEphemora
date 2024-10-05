// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	models "EphemoraApi/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserService) AddUser(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserServiceMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserService)(nil).AddUser), user)
}

// Login mocks base method.
func (m *MockUserService) Login(user models.UserDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), user)
}

// MockLeaderboardService is a mock of LeaderboardService interface.
type MockLeaderboardService struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderboardServiceMockRecorder
}

// MockLeaderboardServiceMockRecorder is the mock recorder for MockLeaderboardService.
type MockLeaderboardServiceMockRecorder struct {
	mock *MockLeaderboardService
}

// NewMockLeaderboardService creates a new mock instance.
func NewMockLeaderboardService(ctrl *gomock.Controller) *MockLeaderboardService {
	mock := &MockLeaderboardService{ctrl: ctrl}
	mock.recorder = &MockLeaderboardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderboardService) EXPECT() *MockLeaderboardServiceMockRecorder {
	return m.recorder
}

// GetLeaderboard mocks base method.
func (m *MockLeaderboardService) GetLeaderboard() ([]models.LeaderboardEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderboard")
	ret0, _ := ret[0].([]models.LeaderboardEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeaderboard indicates an expected call of GetLeaderboard.
func (mr *MockLeaderboardServiceMockRecorder) GetLeaderboard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderboard", reflect.TypeOf((*MockLeaderboardService)(nil).GetLeaderboard))
}

// UpdateRecord mocks base method.
func (m *MockLeaderboardService) UpdateRecord(recordDTO models.RecordDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", recordDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord.
func (mr *MockLeaderboardServiceMockRecorder) UpdateRecord(recordDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockLeaderboardService)(nil).UpdateRecord), recordDTO)
}
