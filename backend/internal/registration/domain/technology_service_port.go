package domain

import "context"

type TechnologyServicePort interface {
	Get(ctx context.Context) ([]*Technology, error)
	GetByIDs(ctx context.Context, ids []int) ([]*Technology, error)
	Save(ctx context.Context, technology *Technology) (*Technology, error)
}
