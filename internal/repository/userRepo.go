package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	tsql = "postgres"
	url  = "user=newuser dbname=postgres password=Sampishet1 host=localhost sslmode=disable"
)

func InsertUser(user models.User) error {

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

	query := "INSERT INTO ephemora.users(email, password, nickname) VALUES ($1 , $2, $3)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Email, user.Password, user.Nickname)
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

	query := "SELECT COUNT(*) FROM ephemora.users WHERE email = $1"
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
