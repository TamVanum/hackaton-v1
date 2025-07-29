package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	registrationRepository domain.RegistrationRepositoryPort
	roleService            *RoleService
	technologyService      *TechnologyService
}

func NewRegistrationService(registrationRepository domain.RegistrationRepositoryPort, roleService *RoleService, technologyService *TechnologyService) *RegistrationService {
	return &RegistrationService{
		registrationRepository: registrationRepository,
		roleService:            roleService,
		technologyService:      technologyService,
	}
}

func (s *RegistrationService) Save(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*domain.Registration, error) {
	registration, err := domain.NewRegistration(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	// roles, err := s.roleService.Get(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// technologies, err := s.technologyService.Get(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	return s.registrationRepository.Save(ctx, registration)
}
