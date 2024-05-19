package models

import "time"

type Record struct {
	Email            string    `json:"email"`
	Record           int       `json:"record"`
	RegistrationDate time.Time `json:"RegistrationDate"`
	UpdateDate       time.Time `json:"UpdateDate"`
}
