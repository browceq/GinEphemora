package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type userRepo struct {
	driver, url string
}

func NewUserRepo(driver, url string) UserRepo {
	return &userRepo{driver, url}
}

func (uRepo *userRepo) InsertUser(user models.User, record models.Record) (returnErr error) {
	if ins := uRepo.isInserted(user.Email); ins != nil {
		return ins
	}

	db, err := sql.Open(uRepo.driver, uRepo.url)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				returnErr = fmt.Errorf("transaction rolled back due to panic: %v; rollback error: %v", p, txErr)
				return
			}
			returnErr = fmt.Errorf("transaction rolled back due to panic: %v", p)
			return
		}

		if returnErr != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				returnErr = fmt.Errorf("transaction rolled back due to error: %v; rollback error: %v", returnErr, txErr)
				return
			}
			returnErr = fmt.Errorf("transaction rolled back due to error: %v", returnErr)
			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			returnErr = fmt.Errorf("transaction commit error: %v", commitErr)
		}
	}()

	userQuery := "INSERT INTO ephemora.users(email, password, nickname) VALUES ($1, $2, $3)"
	_, returnErr = tx.Exec(userQuery, user.Email, user.Password, user.Nickname)
	if returnErr != nil {
		return
	}

	recordQuery := "INSERT INTO ephemora.leaderboard(email, record, registration_date, update_date) VALUES ($1, $2, $3, $4)"
	_, returnErr = tx.Exec(recordQuery, record.Email, record.Record, record.RegistrationDate, record.UpdateDate)
	if returnErr != nil {
		return
	}

	return
}

func (uRepo *userRepo) isInserted(email string) error {
	db, err := sql.Open(uRepo.driver, uRepo.url)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}

	query := "SELECT * FROM ephemora.users WHERE email = $1"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(email)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 0 {
		return errors.New("email занят")
	}
	return nil
}

func (uRepo *userRepo) Login(email string) (string, error) {

	db, err := sql.Open(uRepo.driver, uRepo.url)
	if err != nil {
		return "", err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return "", err
	}

	query := "SELECT password FROM ephemora.users WHERE email = $1"
	var hashedPassword string
	err = db.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		return "", errors.New("wrong login or password")
	}
	return hashedPassword, nil
}
