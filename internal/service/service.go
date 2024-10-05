package service

import "EphemoraApi/internal/models"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type UserService interface {
	AddUser(user models.User) error
	Login(user models.UserDTO) error
}

type LeaderboardService interface {
	UpdateRecord(recordDTO models.RecordDTO) error
	GetLeaderboard() ([]models.LeaderboardEntry, error)
}
