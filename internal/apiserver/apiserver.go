package apiserver

import (
	repo "github.com/diphantxm/hospital-rest-api/internal/storage/repositories"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type api struct {
	storage *repo.Storage // Database
	config  *Config       // Configuration
	logger  *log.Logger   // Logger
	handler *gin.Engine
}

func NewApiServer(config *Config) *api {
	return &api{
		config:  config,
		storage: repo.NewStorage(config.StorageConfig),
		logger:  log.New(),
		handler: gin.New(),
	}
}

func (a *api) Configure() error {
	if err := a.ConfigureLogger(); err != nil {
		log.Fatal("Logger was not been able to configure.", err)
	}

	if err := a.ConfigureRoutes(); err != nil {
		a.logger.Fatal("Routes were not been able to configure.", err)
	}

	return nil
}

func (a *api) ConfigureLogger() error {
	level, err := log.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)
	a.logger.Formatter.(*log.TextFormatter).DisableColors = false
	a.logger.Formatter.(*log.TextFormatter).TimestampFormat = "2006-01-02 15:04:05"
	a.logger.Info("API server has been configured.")

	return nil
}

func (a *api) ConfigureRoutes() error {
	apiRoute := a.handler.Group("/api")
	{
		v1 := apiRoute.Group("/v1")
		{
			doctors := v1.Group("/doctors")
			{
				doctors.GET("/", a.GetAllDoctors)
				doctors.GET("/:id", a.GetDoctorById)
				doctors.GET("/specialties/:specialty", a.GetAllDoctorsBySpecialty)
				doctors.DELETE("/:id", a.DeleteDoctorById)
			}

			patients := v1.Group("/patients")
			{
				patients.GET("/", a.GetAllPatients)
				patients.GET("/:id", a.GetPatientById)
				patients.DELETE("/:id", a.DeletePatientById)

			}

			diseases := v1.Group("/diseases")
			{
				diseases.GET("/", a.GetAllDiseases)
				diseases.GET("/:id", a.GetDiseaseById)
				diseases.GET("/patients/:id", a.GetAllDiseasesByPatientId)
				diseases.GET("/diseases/:diseaseName", a.GetAllDiseasesByName) // -
				diseases.DELETE("/:id", a.DeleteDiseaseById)
			}

			visits := v1.Group("/visits")
			{
				visits.GET("/", a.GetAllVisists)
				visits.GET("/:id", a.GetVisitById)
				visits.GET("/patients/:id", a.GetAllVisitsByPatientId)
				visits.GET("/diseases/:id", a.GetAllVisitsByDiseaseId)
				visits.GET("/doctors/:id", a.GetAllVisitsByDoctorId)
				visits.GET("/date/:date", a.GetAllVisitsByDate) // -
				visits.DELETE("/:id", a.DeleteVisitById)
			}
		}
	}

	return nil
}

func (a *api) Run() error {
	if a.storage == nil {
		a.logger.Fatal("Database instance is nil.")
	}

	if err := a.storage.Open(); err != nil {
		a.logger.Fatal("Error occurred while opening database.", err)
		return err
	}

	a.logger.Info("API server is running...")
	a.handler.Run(a.config.Port)

	return nil
}

func (a *api) Stop() error {
	if err := a.storage.Close(); err != nil {
		a.logger.Fatal("Error occurred while closing database.", err)
		return err
	}

	a.logger.Info("API server has been stopped.")

	return nil
}
