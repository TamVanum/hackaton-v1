package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/app"
	repo "hackathon-pvc-backend/internal/registration/infrastructure/persistance/sql"
	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	RegistrationHandler *rest.RegistrationHandler
}

func WireDependencies(db *sql.DB) *Dependencies {
	registrationRepository := repo.NewSqlRegistrationRepository(db)
	roleRepository := repo.NewSqlRoleRepository(db)
	technologyRepository := repo.NewSqlTechnologyRepository(db)

	roleService := app.NewRoleService(roleRepository)
	technologyService := app.NewTechnologyService(technologyRepository)
	registrationService := app.NewRegistrationService(registrationRepository, roleService, technologyService)

	registrationHandler := rest.NewRegistrationHandler(registrationService)

	return &Dependencies{
		RegistrationHandler: registrationHandler,
	}
}
