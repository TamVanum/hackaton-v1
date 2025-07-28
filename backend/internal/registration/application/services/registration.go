package services

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	registrationRepository domain.RegistrationRepositoryPort
}

func NewRegistrationService(registrationRepository domain.RegistrationRepositoryPort) *RegistrationService {
	return &RegistrationService{
		registrationRepository: registrationRepository,
	}
}

func (s *RegistrationService) Save(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*domain.Registration, error) {
	registration, err := domain.NewRegistration(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	return s.registrationRepository.Save(ctx, registration)
}
