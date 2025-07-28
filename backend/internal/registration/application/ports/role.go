package ports

import (
	"context"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RoleRepositoryPort interface {
	Save(ctx context.Context, role *domain.Role) (*domain.Role, error)
}
