package domain

import (
	"context"
)

type RegistrationRepositoryPort interface {
	Save(ctx context.Context, registration *Registration) (*Registration, error)
}
