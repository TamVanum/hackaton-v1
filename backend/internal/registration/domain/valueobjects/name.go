package valueobjects

import (
	"errors"
	"strings"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, errors.New("name is required")
	}

	if len(value) < 2 {
		return nil, errors.New("name must be at least 2 characters")
	}

	if len(value) > 100 {
		return nil, errors.New("name must be less than 100 characters")
	}

	return &Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}

func (n *Name) String() string {
	return n.value
}
