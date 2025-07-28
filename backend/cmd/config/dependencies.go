package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/application/services"
	repo "hackathon-pvc-backend/internal/registration/infrastructure/persistance/sql"
	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	RegistrationHandler *rest.RegistrationHandler
}

func WireDependencies(db *sql.DB) *Dependencies {
	repository := repo.NewSqlRegistrationRepository(db)

	registrationService := services.NewRegistrationService(repository)

	registrationHandler := rest.NewRegistrationHandler(registrationService)

	return &Dependencies{
		RegistrationHandler: registrationHandler,
	}
}
