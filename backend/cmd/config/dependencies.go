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
	registrationRepository := repo.NewSqlRegistrationRepository(db)
	// roleRepository := repo.NewSqlRoleRepository(db)
	// technologyRepository := repo.NewSqlTechnologyRepository(db)

	registrationService := services.NewRegistrationService(registrationRepository)
	// roleService := services.NewRoleService(roleRepository)
	// technologyService := services.NewTechnologyService(technologyRepository)

	registrationHandler := rest.NewRegistrationHandler(registrationService)

	return &Dependencies{
		RegistrationHandler: registrationHandler,
	}
}
