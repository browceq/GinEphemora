package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
)

func AddUser(user models.User) error {
	err := repository.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}
