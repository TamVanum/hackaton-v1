package valueobjects

import (
	"errors"
	"strings"
)

type ProjectIdea struct {
	value string
}

func NewProjectIdea(value string) (*ProjectIdea, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, errors.New("project idea is required")
	}

	if len(value) < 10 {
		return nil, errors.New("project idea must be at least 10 characters")
	}

	if len(value) > 1000 {
		return nil, errors.New("project idea must be less than 1000 characters")
	}

	return &ProjectIdea{value: value}, nil
}

func (p *ProjectIdea) Value() string {
	return p.value
}

func (p *ProjectIdea) String() string {
	return p.value
}
