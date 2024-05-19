package services

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"fmt"
	"time"
)

func AddUser(user models.User) error {

	record := models.Record{
		Email:            user.Email,
		Record:           0,
		RegistrationDate: time.Now().UTC(),
		UpdateDate:       time.Now().UTC(),
	}

	err := repository.InsertUser(user, record)
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
