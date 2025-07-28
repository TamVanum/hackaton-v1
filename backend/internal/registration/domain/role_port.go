package domain

import (
	"context"
)

type RoleRepositoryPort interface {
	Save(ctx context.Context, role *Role) (*Role, error)
	FindAll(ctx context.Context) ([]*Role, error)
	BulkFindByIDs(ctx context.Context, ids []int) ([]*Role, error)
}
