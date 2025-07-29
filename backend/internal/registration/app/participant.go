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
