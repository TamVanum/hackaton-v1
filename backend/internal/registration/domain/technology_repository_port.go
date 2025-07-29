package domain

import "context"

// TechnologyRepositoryPort defines the contract for technology persistence
type TechnologyRepositoryPort interface {
	FindAll(ctx context.Context) ([]*Technology, error)
	FindByIDs(ctx context.Context, ids []int) ([]*Technology, error)
	Save(ctx context.Context, technology *Technology) (*Technology, error)
}
