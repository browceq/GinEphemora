package repository

import "EphemoraApi/internal/models"

type UserRepo interface {
	InsertUser(user models.User, record models.Record) (returnErr error)
	isInserted(email string) error
	Login(email string) (string, error)
}

type LeaderboardRepo interface {
	UpdateRecord(record models.Record) error
	SelectLeaderboard() ([]models.LeaderboardEntry, error)
}
