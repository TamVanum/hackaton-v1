package services

import (
	"context"
	"hackathon-pvc-backend/internal/registration/application/ports"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	repository ports.RepositoryPort
}

func NewRegistrationService(repository ports.RepositoryPort) *RegistrationService {
	return &RegistrationService{
		repository: repository,
	}
}

func (s *RegistrationService) RegisterParticipant(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*domain.Registration, error) {
	registration, err := domain.NewRegistration(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	return s.repository.Save(ctx, registration)
}
