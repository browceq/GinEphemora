package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"fmt"
)

func AddUser(user models.User) error {
	err := repository.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Login(user models.UserDTO) error {
	err := repository.Login(user)
	if err != nil {
		return err
	}
	return nil
}
