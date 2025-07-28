package ports

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationRepositoryPort interface {
	Save(ctx context.Context, registration *domain.Registration) (*domain.Registration, error)
}
