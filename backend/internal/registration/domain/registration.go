package domain

import (
	"hackathon-pvc-backend/internal/registration/domain/valueobjects"
)

// Registration is the aggregate root
type Registration struct {
	id          *valueobjects.RegistrationID
	name        *valueobjects.Name
	nickname    *valueobjects.Nickname
	projectIdea *valueobjects.ProjectIdea
	teammate    *valueobjects.Teammate
	role        *valueobjects.Role
	createdAt   *valueobjects.CreatedAt
}

func NewRegistration(name, nickname, projectIdea string, teammate *string, role string) (*Registration, error) {
	nameVO, err := valueobjects.NewName(name)
	if err != nil {
		return nil, err
	}

	nicknameVO, err := valueobjects.NewNickname(nickname)
	if err != nil {
		return nil, err
	}

	projectIdeaVO, err := valueobjects.NewProjectIdea(projectIdea)
	if err != nil {
		return nil, err
	}

	teammateVO, err := valueobjects.NewTeammate(teammate)
	if err != nil {
		return nil, err
	}

	roleVO, err := valueobjects.NewRole(role)
	if err != nil {
		return nil, err
	}

	return &Registration{
		name:        nameVO,
		nickname:    nicknameVO,
		projectIdea: projectIdeaVO,
		teammate:    teammateVO,
		role:        roleVO,
		createdAt:   valueobjects.NewCreatedAtNow(),
	}, nil
}

// Getters
func (r *Registration) ID() *valueobjects.RegistrationID {
	return r.id
}

func (r *Registration) Name() *valueobjects.Name {
	return r.name
}

func (r *Registration) Nickname() *valueobjects.Nickname {
	return r.nickname
}

func (r *Registration) ProjectIdea() *valueobjects.ProjectIdea {
	return r.projectIdea
}

func (r *Registration) Teammate() *valueobjects.Teammate {
	return r.teammate
}

func (r *Registration) Role() *valueobjects.Role {
	return r.role
}

func (r *Registration) CreatedAt() *valueobjects.CreatedAt {
	return r.createdAt
}

// SetID is used by repository after persistence
func (r *Registration) SetID(id int) error {
	idVO, err := valueobjects.NewRegistrationID(id)
	if err != nil {
		return err
	}
	r.id = idVO
	return nil
}
