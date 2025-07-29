package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/app"
	repo "hackathon-pvc-backend/internal/registration/infrastructure/persistance/sql"
	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	ParticipantHandler *rest.ParticipantHandler
}

func WireDependencies(db *sql.DB) *Dependencies {
	participantRepository := repo.NewSqlParticipantRepository(db)
	// roleRepository := repo.NewSqlRoleRepository(db)
	// technologyRepository := repo.NewSqlTechnologyRepository(db)

	participantService := app.NewParticipantService(participantRepository)

	// roleService := services.NewRoleService(roleRepository)
	// technologyService := services.NewTechnologyService(technologyRepository)

	participantHandler := rest.NewParticipantHandler(participantService)

	return &Dependencies{
		ParticipantHandler: participantHandler,
	}
}
