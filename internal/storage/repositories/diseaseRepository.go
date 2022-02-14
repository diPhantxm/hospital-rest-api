package repositories

import (
	"database/sql"

	"github.com/diphantxm/hospital-rest-api/internal/storage/models"
)

type DiseaseRepository struct {
	storage *Storage
}

func (repo *DiseaseRepository) GetAll() ([]models.Disease, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[diseases] 
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id`,
	)

	if err != nil {
		return nil, err
	}

	var diseases []models.Disease
	for rows.Next() {
		disease, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		diseases = append(diseases, *disease)
	}

	return diseases, nil
}

func (repo *DiseaseRepository) Add(disease *models.Disease) (*models.Disease, error) {
	if err := repo.storage.db.QueryRow(
		"INSERT INTO [hospital-rest-api].[dbo].[diseases] (diseaseName, treatment, startDate, patientId) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3, @p4)",
		disease.Name,
		disease.Treatment,
		disease.Date.Format("2006-01-02"),
		disease.PatientId,
	).Scan(&disease.Id); err != nil {
		return disease, err
	}

	return disease, nil
}

func (repo *DiseaseRepository) Remove(id int) error {
	if err := repo.storage.db.QueryRow(
		"DELETE FROM [hospital-rest-api].[dbo].[diseases] WHERE id=@p1",
		id,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *DiseaseRepository) FindById(id int) (*models.Disease, error) {
	row := repo.storage.db.QueryRow(
		`SELECT * FROM [hospital-rest-api].[dbo].[diseases] 
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		WHERE diseases.id=@p1`,
		id,
	)

	disease, err := repo.FillFromRow(row)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return disease, nil
}

func (repo *DiseaseRepository) FindAllByPatientId(id int) ([]models.Disease, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[diseases]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		WHERE patientId=@p1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var diseases []models.Disease
	for rows.Next() {
		disease, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		diseases = append(diseases, *disease)
	}

	return diseases, nil
}

func (repo *DiseaseRepository) FindAllByName(name string) ([]models.Disease, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[diseases] 
		LEFT JOIN patients ON patientId=patients.Id 
		WHERE diseaseName=@p1`,
		name,
	)

	if err != nil {
		return nil, err
	}

	var diseases []models.Disease
	for rows.Next() {
		disease, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		diseases = append(diseases, *disease)
	}

	return diseases, nil
}

func (repo *DiseaseRepository) Edit(disease *models.Disease) (*models.Disease, error) {
	row := repo.storage.db.QueryRow(
		`UPDATE [hospital-rest-api].[dbo].[diseases]
		SET diseaseName=@p1, treatment=@p2, startDate=@p3, patientId=@p4, discharged=@p5
		WHERE id=@p6`,
		disease.Name, disease.Treatment, disease.Date, disease.PatientId, disease.Discharged, disease.Id,
	)

	newDisease, err := repo.FillFromRow(row)
	if err != nil {
		return disease, err
	}

	return newDisease, nil
}

func (repo *DiseaseRepository) FillFromRow(diseaseRow *sql.Row) (*models.Disease, error) {
	disease := models.NewDisease()
	if err := diseaseRow.Scan(&disease.Id, &disease.Name, &disease.Treatment, &disease.Date,
		&disease.PatientId, &disease.Patient.Id, &disease.Patient.FirstName,
		&disease.Patient.LastName, &disease.Patient.BirthDate, &disease.Patient.Residence); err != nil {
		return nil, err
	}

	return disease, nil
}

func (repo *DiseaseRepository) FillFromRows(diseaseRow *sql.Rows) (*models.Disease, error) {
	disease := models.NewDisease()
	if err := diseaseRow.Scan(&disease.Id, &disease.Name, &disease.Treatment, &disease.Date,
		&disease.PatientId, &disease.Patient.Id, &disease.Patient.FirstName,
		&disease.Patient.LastName, &disease.Patient.BirthDate, &disease.Patient.Residence); err != nil {
		return nil, err
	}

	return disease, nil
}
