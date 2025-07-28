package domain

import (
	"context"
)

type TechnologyRepositoryPort interface {
	Save(ctx context.Context, technology *Technology) (*Technology, error)
	FindAll(ctx context.Context) ([]*Technology, error)
	BulkFindByIDs(ctx context.Context, ids []int) ([]*Technology, error)
}
