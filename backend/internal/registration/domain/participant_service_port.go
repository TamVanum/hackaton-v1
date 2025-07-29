package domain

import "context"

type ParticipantServicePort interface {
	RegisterParticipant(ctx context.Context, participant *Participant) (*Participant, error)
	FindByNickname(ctx context.Context, nickname string) (*Participant, error)
	AssignRoles(ctx context.Context, participant *Participant, roles []*Role) error
	AssignTechnologies(ctx context.Context, participant *Participant, technologies []*Technology) error
}
