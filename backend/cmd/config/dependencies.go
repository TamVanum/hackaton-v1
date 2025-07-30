package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	RegistrationHandler *rest.RegistrationHandler
}

func WireDependencies(db *sql.DB) *Dependencies {

	return &Dependencies{
		RegistrationHandler: nil,
	}
}
