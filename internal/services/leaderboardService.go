package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"time"
)

func UpdateRecord(recordDTO models.RecordDTO) error {
	record := models.Record{
		Email:      recordDTO.Email,
		Record:     recordDTO.Record,
		UpdateDate: time.Now().UTC(),
	}

	err := repository.UpdateRecord(record)
	return err
}

func GetLeaderboard() ([]models.LeaderboardEntry, error) {
	leaderboard, err := repository.SelectLeaderboard()
	return leaderboard, err
}
