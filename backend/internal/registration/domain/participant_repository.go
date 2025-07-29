package domain

import "context"

type ParticipantRepositoryPort interface {
	Save(ctx context.Context, participant *Participant) (*Participant, error)
	FindByID(ctx context.Context, id int) (*Participant, error)
	FindByNickname(ctx context.Context, nickname string) (*Participant, error)
	FindAll(ctx context.Context) ([]*Participant, error)
	SetRoles(ctx context.Context, participant *Participant, roles []*Role) error
	SetTechnologies(ctx context.Context, participant *Participant, technologies []*Technology) error
}
