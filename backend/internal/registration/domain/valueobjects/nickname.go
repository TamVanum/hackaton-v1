package valueobjects

import (
	"errors"
	"strings"
)

type Nickname struct {
	value string
}

func NewNickname(value string) (*Nickname, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, errors.New("nickname is required")
	}

	if len(value) < 2 {
		return nil, errors.New("nickname must be at least 2 characters")
	}

	if len(value) > 50 {
		return nil, errors.New("nickname must be less than 50 characters")
	}

	return &Nickname{value: value}, nil
}

func (n *Nickname) Value() string {
	return n.value
}

func (n *Nickname) String() string {
	return n.value
}
