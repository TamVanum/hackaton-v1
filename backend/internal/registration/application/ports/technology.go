package ports

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type TechnologyRepositoryPort interface {
	Save(ctx context.Context, technology *domain.Technology) (*domain.Technology, error)
}
