package domain

import (
	"context"
)

type RoleRepositoryPort interface {
	Save(ctx context.Context, role *Role) (*Role, error)
}
