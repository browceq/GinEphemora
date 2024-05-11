package models

import "time"

type Record struct {
	Email            string
	Record           int
	RegistrationDate time.Time
	UpdateDate       time.Time
}
