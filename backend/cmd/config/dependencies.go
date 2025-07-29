package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	ParticipantHandler *rest.ParticipantHandler
}

func WireDependencies(db *sql.DB) *Dependencies {

	return &Dependencies{
		ParticipantHandler: nil,
	}
}
