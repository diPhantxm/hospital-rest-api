package models

import "time"

type Disease struct {
	Id        int
	Name      string
	Treatment string
	Date      time.Time
	PatientId int
	Patient   *Patient
}

func NewDisease() *Disease {
	return &Disease{
		Patient: NewPatient(),
	}
}
