package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"
	"errors"
)

func InsertRecord(record models.Record) error {
	db, err := sql.Open(tsql, url)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	query := "INSERT INTO ephemora.leaderboard(email, record, registrationDate, updateDate) VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(record.Email, record.Record, record.RegistrationDate, record.UpdateDate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(record models.Record) error {
	db, err := sql.Open(tsql, url)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	query := "Update ephemora.leaderboard SET record = $1, updateDate = $2 WHERE email = $3"
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
		return errors.New("ошибка обновления")
	}
	return nil
}

func SelectLeaderboard() ([]models.LeaderboardEntry, error) {
	db, err := sql.Open(tsql, url)
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
