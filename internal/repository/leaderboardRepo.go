package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"
	"errors"
)

type LeaderboardRepo interface {
	UpdateRecord(record models.Record) error
	SelectLeaderboard() ([]models.LeaderboardEntry, error)
}

type leaderboarRepo struct {
	driver, url string
}

func NewLeaderboardRepo(driver, url string) LeaderboardRepo {
	return &leaderboarRepo{driver, url}
}

func (lrRepo *leaderboarRepo) UpdateRecord(record models.Record) error {
	db, err := sql.Open(lrRepo.driver, lrRepo.url)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	query := "Update ephemora.leaderboard SET record = $1, update_date = $2 WHERE email = $3 AND update_date <> $2 AND record <= (CURRENT_DATE - registration_date)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(record.Record, record.UpdateDate, record.Email)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("failed to update record")
	}
	return nil
}

func (lrRepo *leaderboarRepo) SelectLeaderboard() ([]models.LeaderboardEntry, error) {
	db, err := sql.Open(lrRepo.driver, lrRepo.url)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	query := "SELECT ephemora.users.nickname, ephemora.leaderboard.record FROM ephemora.users JOIN ephemora.leaderboard ON ephemora.users.email = ephemora.leaderboard.email ORDER BY ephemora.leaderboard.record DESC LIMIT 750"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var leaderboard []models.LeaderboardEntry
	for rows.Next() {
		var entry models.LeaderboardEntry
		if err := rows.Scan(&entry.Nickname, &entry.Record); err != nil {
			return nil, err
		}
		leaderboard = append(leaderboard, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return leaderboard, nil

}
