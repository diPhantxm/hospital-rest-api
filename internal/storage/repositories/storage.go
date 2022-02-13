package repositories

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/diphantxm/hospital-rest-api/internal/storage"
	//_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
)

type Storage struct {
	diseaseRepo *DiseaseRepository
	doctorRepo  *DoctorRepository
	visitRepo   *VisitRepository
	patientRepo *PatientRepository
	db          *sql.DB
	config      *storage.Config
}

func NewStorage(config *storage.Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("sqlserver", s.config.DataSourceName)

	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Storage) Close() error {

	if s.db == nil {
		return nil
	}

	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Doctor() *DoctorRepository {
	if s.doctorRepo != nil {
		return s.doctorRepo
	}

	s.doctorRepo = &DoctorRepository{
		storage: s,
	}

	return s.doctorRepo
}

func (s *Storage) Disease() *DiseaseRepository {
	if s.diseaseRepo != nil {
		return s.diseaseRepo
	}

	s.diseaseRepo = &DiseaseRepository{
		storage: s,
	}

	return s.diseaseRepo
}

func (s *Storage) Patient() *PatientRepository {
	if s.patientRepo != nil {
		return s.patientRepo
	}

	s.patientRepo = &PatientRepository{
		storage: s,
	}

	return s.patientRepo
}

func (s *Storage) Visit() *VisitRepository {
	if s.visitRepo != nil {
		return s.visitRepo
	}

	s.visitRepo = &VisitRepository{
		storage: s,
	}

	return s.visitRepo
}
