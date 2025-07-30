package app

import (
	"context"
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

func (s *RegistrationService) Create(
	ctx context.Context,
	name, nickname, email, region, projectIdea string,
	teamPreference bool,
	desiredTeammate *string,
	roleIDs []int,
	technologyIDs []int,
) (*domain.Registration, error) {

	if err := s.participantService.CheckNicknameAvailability(ctx, nickname); err != nil {
		return nil, err
	}

	participant, err := s.participantService.Make(ctx, name, nickname, email, region, projectIdea, teamPreference, desiredTeammate)
	if err != nil {
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

	return domain.NewRegistration(participant, roles, technologies)
}

func (s *RegistrationService) Persist(ctx context.Context, registration *domain.Registration) (*domain.Registration, error) {

	savedParticipant, err := s.participantService.Persist(ctx, registration.Participant())
	if err != nil {
		return nil, err
	}

	if err := s.participantService.AssignRoles(ctx, savedParticipant, registration.Roles()); err != nil {
		return nil, err
	}

	if err := s.participantService.AssignTechnologies(ctx, savedParticipant, registration.Technologies()); err != nil {
		return nil, err
	}

	completedRegistration, err := domain.NewRegistration(savedParticipant, registration.Roles(), registration.Technologies())
	if err != nil {
		return nil, err
	}

	completedRegistration.SetID(savedParticipant.ID())
	completedRegistration.SetCreatedAt(savedParticipant.CreatedAt())

	return completedRegistration, nil
}

func (s *RegistrationService) Register(
	ctx context.Context,
	name, nickname, email, region, projectIdea string,
	teamPreference bool,
	desiredTeammate *string,
	roleIDs []int,
	technologyIDs []int,
) (*domain.Registration, error) {

	registration, err := s.Create(ctx, name, nickname, email, region, projectIdea, teamPreference, desiredTeammate, roleIDs, technologyIDs)
	if err != nil {
		return nil, err
	}

	return s.Persist(ctx, registration)
}
