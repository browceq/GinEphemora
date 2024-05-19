package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	tsql = "postgres"
	url  = "user=newuser dbname=postgres password=Sampishet1 host=localhost sslmode=disable"
)

func InsertUser(user models.User, record models.Record) error {

	if ins := isInserted(user.Email); ins != nil {
		return ins
	}

	db, err := sql.Open(tsql, url)
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
			tx.Rollback()
			err = fmt.Errorf("transaction rolled back due to panic: %v", p)
		} else if err != nil {
			tx.Rollback()
			err = fmt.Errorf("transaction rolled back due to error: %v", err)
		} else {
			err = tx.Commit()
		}
	}()

	userQuery := "INSERT INTO ephemora.users(email, password, nickname) VALUES ($1 , $2, $3)"
	_, err = tx.Exec(userQuery, user.Email, user.Password, user.Nickname)
	if err != nil {
		return err
	}

	recordQuery := "INSERT INTO ephemora.leaderboard(email, record, registration_date, update_date) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(recordQuery, record.Email, record.Record, record.RegistrationDate, record.UpdateDate)

	if err != nil {
		return err
	}

	return nil
}

func isInserted(email string) error {
	db, err := sql.Open(tsql, url)
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

func Login(user models.UserDTO) error {

	db, err := sql.Open(tsql, url)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}

	query := "SELECT count(*) FROM ephemora.users WHERE email = $1 and password = $2"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("неправильный логин или пароль")
	}
	return nil
}
