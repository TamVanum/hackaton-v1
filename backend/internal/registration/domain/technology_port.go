package domain

import (
	"context"
)

type TechnologyRepositoryPort interface {
	Save(ctx context.Context, technology *Technology) (*Technology, error)
}
