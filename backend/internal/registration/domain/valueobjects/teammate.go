package valueobjects

import (
	"errors"
	"strings"
)

type Teammate struct {
	value *string
}

func NewTeammate(value *string) (*Teammate, error) {
	if value == nil {
		return &Teammate{value: nil}, nil
	}

	trimmed := strings.TrimSpace(*value)

	if trimmed == "" {
		return &Teammate{value: nil}, nil
	}

	if len(trimmed) < 2 {
		return nil, errors.New("teammate name must be at least 2 characters")
	}

	if len(trimmed) > 100 {
		return nil, errors.New("teammate name must be less than 100 characters")
	}

	return &Teammate{value: &trimmed}, nil
}

func (t *Teammate) Value() *string {
	return t.value
}

func (t *Teammate) HasValue() bool {
	return t.value != nil && *t.value != ""
}

func (t *Teammate) String() string {
	if t.value == nil {
		return ""
	}
	return *t.value
}
