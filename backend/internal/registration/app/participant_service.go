package app

import (
	"context"
	"errors"
	"fmt"
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

func (s *ParticipantService) Make(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*domain.Participant, error) {
	participant, err := domain.NewParticipant(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	return participant, nil
}

func (s *ParticipantService) CheckNicknameAvailability(ctx context.Context, nickname string) error {
	existingParticipant, err := s.participantRepository.FindByNickname(ctx, nickname)
	if err == nil && existingParticipant != nil {
		return errors.New("nickname already taken")
	}

	return nil
}

func (s *ParticipantService) Persist(ctx context.Context, domainParticipant *domain.Participant) (*domain.Participant, error) {
	persistedParticipant, err := s.participantRepository.Save(ctx, domainParticipant)
	if err != nil {
		return nil, err
	}

	fmt.Println("persistedParticipant", persistedParticipant)
	return persistedParticipant, nil
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
