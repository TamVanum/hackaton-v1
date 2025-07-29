package domain

import "context"

type RoleServicePort interface {
	Get(ctx context.Context) ([]*Role, error)
	GetByIDs(ctx context.Context, ids []int) ([]*Role, error)
	Save(ctx context.Context, role *Role) (*Role, error)
}
