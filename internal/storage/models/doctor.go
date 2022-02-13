package models

type Doctor struct {
	Id        int
	FirstName string
	LastName  string
	Specialty string
}

func NewDoctor() *Doctor {
	return &Doctor{}
}
