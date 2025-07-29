package domain

import "context"

// ParticipantRepositoryPort defines the contract for participant persistence
type ParticipantRepositoryPort interface {
	Save(ctx context.Context, participant *Participant) (*Participant, error)
	FindByID(ctx context.Context, id int) (*Participant, error)
	FindByNickname(ctx context.Context, nickname string) (*Participant, error)
	FindAll(ctx context.Context) ([]*Participant, error)
	Count(ctx context.Context) (int, error)
}
