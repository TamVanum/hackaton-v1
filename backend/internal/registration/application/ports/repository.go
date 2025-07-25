package ports

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

// RepositoryPort defines the contract for registration persistence
type RepositoryPort interface {
	Save(ctx context.Context, registration *domain.Registration) (*domain.Registration, error)
	FindByID(ctx context.Context, id int) (*domain.Registration, error)
	FindByNickname(ctx context.Context, nickname string) (*domain.Registration, error)
	FindAll(ctx context.Context) ([]*domain.Registration, error)
	Count(ctx context.Context) (int, error)
}
