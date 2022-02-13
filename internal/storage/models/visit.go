package models

import "time"

type Visit struct {
	Id   int
	Date time.Time

	PatientId int
	Patient   *Patient

	DiseaseId int
	Disease   *Disease

	DoctorId int
	Doctor   *Doctor
}

func NewVisit() *Visit {
	return &Visit{
		Patient: NewPatient(),
		Disease: NewDisease(),
		Doctor:  NewDoctor(),
	}
}
