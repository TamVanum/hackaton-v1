package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type ParticipantService struct {
	participantRepository domain.ParticipantRepository
}

func NewParticipantService(participantRepository domain.ParticipantRepository) domain.ParticipantService {
	return &ParticipantService{
		participantRepository: participantRepository,
	}
}

func (s *ParticipantService) RegisterParticipant(ctx context.Context, participant *domain.Participant) (*domain.Participant, error) {

	return s.participantRepository.Save(ctx, participant)
}

func (s *ParticipantService) FindByNickname(ctx context.Context, nickname string) (*domain.Participant, error) {
	return s.participantRepository.FindByNickname(ctx, nickname)
}

func (s *ParticipantService) AssignRoles(ctx context.Context, participant *domain.Participant, roles []*domain.Role) error {
	if err := s.participantRepository.SetRoles(ctx, participant, roles); err != nil {
		return err
	}

	return nil
}

func (s *ParticipantService) AssignTechnologies(ctx context.Context, participant *domain.Participant, technologies []*domain.Technology) error {
	if err := s.participantRepository.SetTechnologies(ctx, participant, technologies); err != nil {
		return err
	}

	return nil
}
