package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"time"
)

type LeaderboardService interface {
	UpdateRecord(recordDTO models.RecordDTO) error
	GetLeaderboard() ([]models.LeaderboardEntry, error)
}

type leaderboardService struct {
	lrRepo repository.LeaderboardRepo
}

func NewLeaderboardService(repo repository.LeaderboardRepo) LeaderboardService {
	return &leaderboardService{repo}
}

func (lrService *leaderboardService) UpdateRecord(recordDTO models.RecordDTO) error {
	record := models.Record{
		Email:      recordDTO.Email,
		Record:     recordDTO.Record,
		UpdateDate: time.Now().UTC(),
	}

	err := lrService.lrRepo.UpdateRecord(record)
	return err
}

func (lrService *leaderboardService) GetLeaderboard() ([]models.LeaderboardEntry, error) {
	leaderboard, err := lrService.lrRepo.SelectLeaderboard()
	return leaderboard, err
}
