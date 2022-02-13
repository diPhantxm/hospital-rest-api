package repositories

import (
	"database/sql"

	"github.com/diphantxm/hospital-rest-api/internal/storage/models"
)

type DoctorRepository struct {
	storage *Storage
}

func (repo *DoctorRepository) GetAll() ([]models.Doctor, error) {
	rows, err := repo.storage.db.Query(
		"SELECT * FROM [hospital-rest-api].[dbo].[doctors]",
	)

	if err != nil {
		return nil, err
	}

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.Specialty); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (repo *DoctorRepository) Add(doctor *models.Doctor) (*models.Doctor, error) {
	if err := repo.storage.db.QueryRow("INSERT INTO [hospital-rest-api].[dbo].[doctors] (firstName, lastName, specialty) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3)",
		doctor.FirstName,
		doctor.LastName,
		doctor.Specialty,
	).Scan(&doctor.Id); err != nil {
		return doctor, err
	}

	return doctor, nil
}

func (repo *DoctorRepository) Remove(id int) error {
	if err := repo.storage.db.QueryRow(
		"DELETE FROM [hospital-rest-api].[dbo].[doctors] WHERE id=@p1",
		id,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *DoctorRepository) FindById(id int) (*models.Doctor, error) {
	doctor := models.NewDoctor()

	if err := repo.storage.db.QueryRow(
		"SELECT * FROM [hospital-rest-api].[dbo].[doctors] WHERE id=@p1",
		id,
	).Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.Specialty); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return doctor, nil
}

func (repo *DoctorRepository) FindAllBySpecialty(specialty string) ([]models.Doctor, error) {
	rows, err := repo.storage.db.Query(
		"SELECT id, firstName, lastName, specialty FROM [hospital-rest-api].[dbo].[doctors] WHERE specialty=@p1",
		specialty,
	)

	if err != nil {
		return nil, err
	}

	var doctors []models.Doctor
	for rows.Next() {
		var doctor models.Doctor
		if err := rows.Scan(&doctor.Id, &doctor.FirstName, &doctor.LastName, &doctor.Specialty); err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}
