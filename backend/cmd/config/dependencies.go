package config

import (
	"database/sql"

	"hackathon-pvc-backend/internal/registration/app"
	sqlRepo "hackathon-pvc-backend/internal/registration/infrastructure/persistance/sql"
	"hackathon-pvc-backend/internal/registration/infrastructure/rest"
)

type Dependencies struct {
	RegistrationHandler *rest.RegistrationHandler
	RolesHandler        *rest.RolesHandler
}

func WireDependencies(db *sql.DB) *Dependencies {
	participantRepo := sqlRepo.NewSqlParticipantRepository(db)
	roleRepo := sqlRepo.NewSqlRoleRepository(db)
	technologyRepo := sqlRepo.NewSqlTechnologyRepository(db)

	// Create services
	participantService := app.NewParticipantService(participantRepo)
	roleService := app.NewRoleService(roleRepo)
	technologyService := app.NewTechnologyService(technologyRepo)

	// Create application service
	registrationService := app.NewRegistrationService(
		participantService,
		roleService,
		technologyService,
	)

	// Create handler
	registrationHandler := rest.NewRegistrationHandler(registrationService)
	rolesHandler := rest.NewRolesHandler(roleService.(*app.RoleService))

	return &Dependencies{
		RegistrationHandler: registrationHandler,
		RolesHandler:        rolesHandler,
	}
}
