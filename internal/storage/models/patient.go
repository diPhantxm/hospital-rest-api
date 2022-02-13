package models

import "time"

type Patient struct {
	Id        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Residence string
}

func NewPatient() *Patient {
	return &Patient{}
}
