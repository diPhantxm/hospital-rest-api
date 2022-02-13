package repositories

import (
	"database/sql"
	"time"

	"github.com/diphantxm/hospital-rest-api/internal/storage/models"
)

type VisitRepository struct {
	storage *Storage
}

func (repo *VisitRepository) GetAll() ([]models.Visit, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id`,
	)

	if err != nil {
		return nil, err
	}

	var visits []models.Visit
	for rows.Next() {
		visit, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		visits = append(visits, *visit)
	}

	return visits, nil
}

func (repo *VisitRepository) Add(visit *models.Visit) (*models.Visit, error) {
	if err := repo.storage.db.QueryRow(
		"INSERT INTO [hospital-rest-api].[dbo].[visits] (patientId, diseaseId, doctorId, visitDate) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3, @p4)",
		visit.PatientId,
		visit.DiseaseId,
		visit.DoctorId,
		visit.Date,
	).Scan(&visit.Id); err != nil {

		return visit, err
	}

	return visit, nil
}

func (repo *VisitRepository) Remove(id int) error {
	if err := repo.storage.db.QueryRow(
		"DELETE FROM [hospital-rest-api].[dbo].[visits] WHERE id=@p1",
		id,
	).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *VisitRepository) FindById(id int) (*models.Visit, error) {
	row := repo.storage.db.QueryRow(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits] 
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id
		WHERE visits.id=@p1`,
		id,
	)

	visit, err := repo.FillFromRow(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return visit, nil
}

func (repo *VisitRepository) FindAllByPatientId(id int) ([]models.Visit, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id
		WHERE visits.patientId=@p1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var visists []models.Visit
	for rows.Next() {
		visit, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		visists = append(visists, *visit)
	}

	return visists, nil
}

func (repo *VisitRepository) FindAllByDiseaseId(id int) ([]models.Visit, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id
		WHERE diseaseId=@p1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var visists []models.Visit
	for rows.Next() {
		visit, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		visists = append(visists, *visit)
	}

	return visists, nil
}

func (repo *VisitRepository) FindAllByDoctorId(id int) ([]models.Visit, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id
		WHERE doctorId=@p1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	var visists []models.Visit
	for rows.Next() {
		visit, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		visists = append(visists, *visit)
	}

	return visists, nil
}

func (repo *VisitRepository) FindAllByDate(date time.Time) ([]models.Visit, error) {
	rows, err := repo.storage.db.Query(
		`SELECT * FROM [hospital-rest-api].[dbo].[visits]
		LEFT JOIN [hospital-rest-api].[dbo].[patients] ON patientId=patients.id
		LEFT JOIN [hospital-rest-api].[dbo].[diseases] ON diseaseId=diseases.id
		LEFT JOIN [hospital-rest-api].[dbo].[doctors] ON doctorId=doctors.id
		WHERE visitDate='@p1'`,
		date,
	)

	if err != nil {
		return nil, err
	}

	var visists []models.Visit
	for rows.Next() {
		visit, err := repo.FillFromRows(rows)
		if err != nil {
			return nil, err
		}
		visists = append(visists, *visit)
	}

	return visists, nil
}

func (repo *VisitRepository) FillFromRow(row *sql.Row) (*models.Visit, error) {
	visit := models.NewVisit()
	visit.Disease.Patient = nil
	row.Scan(&visit.Id, &visit.PatientId, &visit.DiseaseId, &visit.DoctorId, &visit.Date,
		&visit.Patient.Id, &visit.Patient.FirstName, &visit.Patient.LastName, &visit.Patient.BirthDate, &visit.Patient.Residence,
		&visit.Disease.Id, &visit.Disease.Name, &visit.Disease.Treatment, &visit.Disease.Date, &visit.Disease.PatientId,
		&visit.Doctor.Id, &visit.Doctor.FirstName, &visit.Doctor.LastName, &visit.Doctor.Specialty)

	return visit, nil
}

func (repo *VisitRepository) FillFromRows(rows *sql.Rows) (*models.Visit, error) {
	visit := models.NewVisit()
	visit.Disease.Patient = nil
	rows.Scan(&visit.Id, &visit.PatientId, &visit.DiseaseId, &visit.DoctorId, &visit.Date,
		&visit.Patient.Id, &visit.Patient.FirstName, &visit.Patient.LastName, &visit.Patient.BirthDate, &visit.Patient.Residence,
		&visit.Disease.Id, &visit.Disease.Name, &visit.Disease.Treatment, &visit.Disease.Date, &visit.Disease.PatientId,
		&visit.Doctor.Id, &visit.Doctor.FirstName, &visit.Doctor.LastName, &visit.Doctor.Specialty)

	return visit, nil
}
