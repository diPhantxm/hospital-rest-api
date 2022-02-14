package models

import "time"

type Disease struct {
	Id         int
	Name       string
	Treatment  string
	Date       time.Time
	Discharged bool

	PatientId int
	Patient   *Patient

	Visit []Visit
}

func NewDisease() *Disease {
	return &Disease{
		Patient: NewPatient(),
		Visit:   []Visit{},
	}
}
