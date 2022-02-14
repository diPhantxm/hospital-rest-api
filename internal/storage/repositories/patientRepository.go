package repositories

import (
	"database/sql"

	"github.com/diphantxm/hospital-rest-api/internal/storage/models"
)

type PatientRepository struct {
	storage *Storage
}

func (repo *PatientRepository) GetAll() ([]models.Patient, error) {
	rows, err := repo.storage.db.Query(
		"SELECT * FROM [hospital-rest-api].[dbo].[patients]",
	)

	if err != nil {
		return nil, err
	}

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		if err := rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.BirthDate, &patient.Residence); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

func (repo *PatientRepository) Add(patient *models.Patient) (*models.Patient, error) {
	if err := repo.storage.db.QueryRow(
		"INSERT INTO [hospital-rest-api].[dbo].[patients] (firstName, lastName, birthDate, residence) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3, @p4)",
		patient.FirstName,
		patient.LastName,
		patient.BirthDate.Format("2006-01-02"),
		patient.Residence,
	).Scan(&patient.Id); err != nil {

		return patient, err
	}

	return patient, nil
}

func (repo *PatientRepository) Remove(id int) error {
	if err := repo.storage.db.QueryRow(
		"DELETE FROM [hospital-rest-api].[dbo].[patients] WHERE id=@p1",
		id,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *PatientRepository) FindById(id int) (*models.Patient, error) {
	patient := models.NewPatient()

	if err := repo.storage.db.QueryRow(
		"SELECT * FROM [hospital-rest-api].[dbo].[patients] WHERE id=@p1",
		id,
	).Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.BirthDate, &patient.Residence); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return patient, nil
}

func (repo *PatientRepository) Edit(patient *models.Patient) (*models.Patient, error) {
	newPatient := models.NewPatient()

	err := repo.storage.db.QueryRow(
		`UPDATE [hospital-rest-api].[dbo].[patients]
		SET firstName=@p1, lastName=@p2, birthDate=@p3, residence=@p4
		WHERE id=@p5`,
		patient.FirstName, patient.LastName, patient.BirthDate, patient.Residence, patient.Id,
	).Scan(&newPatient.Id, &newPatient.FirstName, &newPatient.LastName, &newPatient.BirthDate, &newPatient.Residence)

	if err != nil {
		return patient, err
	}

	return newPatient, nil
}
