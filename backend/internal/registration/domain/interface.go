package domain

import (
	"context"
)

type ParticipantRepository interface {
	Save(ctx context.Context, participant *Participant) (*Participant, error)
	FindByID(ctx context.Context, id int) (*Participant, error)
	FindByNickname(ctx context.Context, nickname string) (*Participant, error)
	FindAll(ctx context.Context) ([]*Participant, error)
	SetRoles(ctx context.Context, participant *Participant, roles []*Role) error
	SetTechnologies(ctx context.Context, participant *Participant, technologies []*Technology) error
}

type ParticipantService interface {
	Create(ctx context.Context, name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*Participant, error)
	Persist(ctx context.Context, participant *Participant) (*Participant, error)
	CheckNicknameAvailability(ctx context.Context, nickname string) error
	AssignRoles(ctx context.Context, participant *Participant, roles []*Role) error
	AssignTechnologies(ctx context.Context, participant *Participant, technologies []*Technology) error
}

type RoleRepository interface {
	FindAll(ctx context.Context) ([]*Role, error)
	FindByIDs(ctx context.Context, ids []int) ([]*Role, error)
	Save(ctx context.Context, role *Role) (*Role, error)
}

type RoleService interface {
	Get(ctx context.Context) ([]*Role, error)
	GetByIDs(ctx context.Context, ids []int) ([]*Role, error)
	Save(ctx context.Context, role *Role) (*Role, error)
}

type TechnologyRepository interface {
	FindAll(ctx context.Context) ([]*Technology, error)
	FindByIDs(ctx context.Context, ids []int) ([]*Technology, error)
	Save(ctx context.Context, technology *Technology) (*Technology, error)
}

type TechnologyService interface {
	Get(ctx context.Context) ([]*Technology, error)
	GetByIDs(ctx context.Context, ids []int) ([]*Technology, error)
	Save(ctx context.Context, technology *Technology) (*Technology, error)
}
