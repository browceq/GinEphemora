package service

import (
	"EphemoraApi/internal/models"
	"EphemoraApi/internal/repository"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{repo}
}

func (service *userService) AddUser(user models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	record := models.Record{
		Email:            user.Email,
		Record:           0,
		RegistrationDate: time.Now().UTC(),
		UpdateDate:       time.Now().UTC(),
	}

	err = service.repo.InsertUser(user, record)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (service *userService) Login(user models.UserDTO) error {
	hashedPassword, err := service.repo.Login(user.Email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return errors.New("wrong login or password")
	}
	return nil
}
