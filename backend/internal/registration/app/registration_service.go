package app

import (
	"context"
	"errors"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	participantService   domain.ParticipantServicePort
	roleRepository       domain.RoleRepositoryPort
	technologyRepository domain.TechnologyRepositoryPort
}

func NewRegistrationService(participantService domain.ParticipantServicePort, roleRepository domain.RoleRepositoryPort, technologyRepository domain.TechnologyRepositoryPort) *RegistrationService {
	return &RegistrationService{
		participantService:   participantService,
		roleRepository:       roleRepository,
		technologyRepository: technologyRepository,
	}
}

func (s *RegistrationService) RegisterParticipant(
	ctx context.Context,
	name, nickname, email, region, projectIdea string,
	teamPreference bool,
	desiredTeammate *string,
	roleIDs []int,
	technologyIDs []int,
) (*domain.Participant, error) {

	if err := s.validateRegistrationRules(roleIDs, technologyIDs); err != nil {
		return nil, err
	}

	if err := s.checkNicknameAvailability(ctx, nickname); err != nil {
		return nil, err
	}

	participant, err := domain.NewParticipant(name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
		return nil, err
	}

	savedParticipant, err := s.participantService.RegisterParticipant(ctx, participant)
	if err != nil {
		return nil, err
	}

	if err := s.createRoleRelationships(ctx, savedParticipant.ID(), roleIDs); err != nil {
		return nil, err
	}

	if err := s.createTechnologyRelationships(ctx, savedParticipant.ID(), technologyIDs); err != nil {
		return nil, err
	}

	return savedParticipant, nil
}

func (s *RegistrationService) validateRegistrationRules(roleIDs, technologyIDs []int) error {
	if len(roleIDs) == 0 {
		return errors.New("at least one role must be selected")
	}

	if len(technologyIDs) == 0 {
		return errors.New("at least one technology must be selected")
	}

	if len(roleIDs) > 3 {
		return errors.New("maximum of 3 roles allowed")
	}

	if len(technologyIDs) > 10 {
		return errors.New("maximum of 10 technologies allowed")
	}

	if s.hasDuplicates(roleIDs) {
		return errors.New("duplicate role IDs are not allowed")
	}

	if s.hasDuplicates(technologyIDs) {
		return errors.New("duplicate technology IDs are not allowed")
	}

	return nil
}

func (s *RegistrationService) checkNicknameAvailability(ctx context.Context, nickname string) error {
	existingParticipant, err := s.participantService.FindByNickname(ctx, nickname)
	if err == nil && existingParticipant != nil {
		return errors.New("nickname already taken")
	}
	return nil
}

func (s *RegistrationService) createRoleRelationships(ctx context.Context, participantID int, roleIDs []int) error {
	return nil
}

func (s *RegistrationService) createTechnologyRelationships(ctx context.Context, participantID int, technologyIDs []int) error {
	return nil
}

func (s *RegistrationService) hasDuplicates(slice []int) bool {
	seen := make(map[int]bool)
	for _, item := range slice {
		if seen[item] {
			return true
		}
		seen[item] = true
	}
	return false
}
