package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/application/services"
	"hackathon-pvc-backend/internal/registration/infrastructure/persistance"
	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	RegistrationHandler *rest.RegistrationHandler
}

func WireDependencies(db *sql.DB) *Dependencies {
	repository := persistance.NewSQLiteRepository(db)

	registrationService := services.NewRegistrationService(repository)

	registrationHandler := rest.NewRegistrationHandler(registrationService)

	return &Dependencies{
		RegistrationHandler: registrationHandler,
	}
}
