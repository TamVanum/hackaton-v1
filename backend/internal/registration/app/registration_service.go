package app

import (
	"context"
	"errors"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationService struct {
	participantService domain.ParticipantService
	roleService        domain.RoleService
	technologyService  domain.TechnologyService
}

func NewRegistrationService(
	participantService domain.ParticipantService,
	roleService domain.RoleService,
	technologyService domain.TechnologyService,
) *RegistrationService {
	return &RegistrationService{
		participantService: participantService,
		roleService:        roleService,
		technologyService:  technologyService,
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

	if err := s.checkNicknameAvailability(ctx, nickname); err != nil {
		return nil, err
	}

	roles, err := s.roleService.GetByIDs(ctx, roleIDs)
	if err != nil {
		return nil, err
	}

	technologies, err := s.technologyService.GetByIDs(ctx, technologyIDs)
	if err != nil {
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

	if err := s.participantService.AssignRoles(ctx, savedParticipant, roles); err != nil {
		return nil, err
	}

	if err := s.participantService.AssignTechnologies(ctx, savedParticipant, technologies); err != nil {
		return nil, err
	}

	return savedParticipant, nil
}

func (s *RegistrationService) checkNicknameAvailability(ctx context.Context, nickname string) error {
	existingParticipant, err := s.participantService.FindByNickname(ctx, nickname)
	if err == nil && existingParticipant != nil {
		return errors.New("nickname already taken")
	}
	return nil
}
