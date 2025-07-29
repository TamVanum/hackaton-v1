package domain

import "context"

// RoleRepositoryPort defines the contract for role persistence
type RoleRepositoryPort interface {
	FindAll(ctx context.Context) ([]*Role, error)
	FindByIDs(ctx context.Context, ids []int) ([]*Role, error)
	Save(ctx context.Context, role *Role) (*Role, error)
}
