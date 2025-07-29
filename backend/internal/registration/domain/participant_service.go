package domain

import "context"

type ParticipantServicePort interface {
	RegisterParticipant(ctx context.Context, participant *Participant) (*Participant, error)
	FindByNickname(ctx context.Context, nickname string) (*Participant, error)
}
