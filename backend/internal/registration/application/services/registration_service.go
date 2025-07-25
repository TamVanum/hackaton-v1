package services

import (
	"context"
	"errors"
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

func (s *RegistrationService) RegisterParticipant(ctx context.Context, name, nickname, projectIdea string, teammate *string, role string) (*domain.Registration, error) {
	_, err := s.repository.FindByNickname(ctx, nickname)
	if err == nil {
		return nil, errors.New("nickname already taken")
	}

	registration, err := domain.NewRegistration(name, nickname, projectIdea, teammate, role)
	if err != nil {
		return nil, err
	}

	return s.repository.Save(ctx, registration)
}

func (s *RegistrationService) GetRegistration(ctx context.Context, id int) (*domain.Registration, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *RegistrationService) GetAllRegistrations(ctx context.Context) ([]*domain.Registration, error) {
	return s.repository.FindAll(ctx)
}

func (s *RegistrationService) GetRegistrationStats(ctx context.Context) (map[string]interface{}, error) {
	count, err := s.repository.Count(ctx)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_registrations": count,
		"available_spots":     500 - count,
	}, nil
}
