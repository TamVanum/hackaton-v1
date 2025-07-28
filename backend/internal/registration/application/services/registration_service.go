package services

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	repository domain.RegistrationRepositoryPort
}

func NewRegistrationService(repository domain.RegistrationRepositoryPort) *RegistrationService {
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
