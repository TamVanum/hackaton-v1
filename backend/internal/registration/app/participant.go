package app

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type ParticipantService struct {
	participantRepository domain.ParticipantRepositoryPort
}

func NewParticipantService(participantRepository domain.ParticipantRepositoryPort) *ParticipantService {
	return &ParticipantService{
		participantRepository: participantRepository,
	}
}

func (s *ParticipantService) Save(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*domain.Participant, error) {
	participant, err := domain.NewParticipant(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	return s.participantRepository.Save(ctx, participant)
}

func (s *ParticipantService) FindByNickname(ctx context.Context, nickname string) (*domain.Participant, error) {
	return s.participantRepository.FindByNickname(ctx, nickname)
}

func (s *ParticipantService) AssignRoles(ctx context.Context, participant *domain.Participant, roles []*domain.Role) error {
	return s.participantRepository.SetRoles(ctx, participant, roles)
}

func (s *ParticipantService) AssignTechnologies(ctx context.Context, participant *domain.Participant, technologies []*domain.Technology) error {
	return s.participantRepository.SetTechnologies(ctx, participant, technologies)
}
