package valueobjects

import (
	"errors"
	"time"
)

type CreatedAt struct {
	value time.Time
}

func NewCreatedAt(value time.Time) (*CreatedAt, error) {
	if value.IsZero() {
		return nil, errors.New("created at cannot be zero")
	}

	return &CreatedAt{value: value}, nil
}

func NewCreatedAtNow() *CreatedAt {
	return &CreatedAt{value: time.Now()}
}

func (c *CreatedAt) Value() time.Time {
	return c.value
}

func (c *CreatedAt) String() string {
	return c.value.Format(time.RFC3339)
}
