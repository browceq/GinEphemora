package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"time"
)

func UpdateRecord(record models.Record) error {
	utcTime := record.UpdateDate.UTC()
	record.UpdateDate = utcTime

	err := repository.UpdateRecord(record)
	return err
}

func AddRecord(email string) error {
	record := models.Record{
		Email:            email,
		Record:           0,
		RegistrationDate: time.Now().UTC(),
		UpdateDate:       time.Now().UTC(),
	}
	err := repository.InsertRecord(record)
	return err
}

func GetLeaderboard() ([]models.LeaderboardEntry, error) {
	leaderboard, err := repository.SelectLeaderboard()
	return leaderboard, err
}
