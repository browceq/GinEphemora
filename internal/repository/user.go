package repository

import (
	"EphemoraApi/internal/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tsql = "postgresql"
	url  = "newuser:Sampishet1@tcp(localhost:5432)/ephemora"
)

func InsertUser(user models.User) error {
	db, err := sql.Open(tsql, url)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	query := "INSERT INTO users(email,password) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
