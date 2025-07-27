package domain

import (
	"errors"
	"strings"
	"time"
)

// Registration is the aggregate root
type Registration struct {
	id              int
	name            string
	nickname        string
	projectIdea     string
	desiredTeammate *string
	role            string
	createdAt       time.Time
}

func NewRegistration(name, nickname, projectIdea string, desiredTeammate *string, role string) (*Registration, error) {
	// Validate name
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	// Validate nickname
	if strings.TrimSpace(nickname) == "" {
		return nil, errors.New("nickname cannot be empty")
	}

	// Validate project idea
	if strings.TrimSpace(projectIdea) == "" {
		return nil, errors.New("project idea cannot be empty")
	}

	// Validate role
	if strings.TrimSpace(role) == "" {
		return nil, errors.New("role cannot be empty")
	}

	return &Registration{
		name:            strings.TrimSpace(name),
		nickname:        strings.TrimSpace(nickname),
		projectIdea:     strings.TrimSpace(projectIdea),
		desiredTeammate: desiredTeammate,
		role:            strings.TrimSpace(role),
		createdAt:       time.Now(),
	}, nil
}

// Getters
func (r *Registration) ID() int {
	return r.id
}

func (r *Registration) Name() string {
	return r.name
}

func (r *Registration) Nickname() string {
	return r.nickname
}

func (r *Registration) ProjectIdea() string {
	return r.projectIdea
}

func (r *Registration) DesiredTeammate() *string {
	return r.desiredTeammate
}

func (r *Registration) Role() string {
	return r.role
}

func (r *Registration) CreatedAt() time.Time {
	return r.createdAt
}

// SetID is used by repository after persistence
func (r *Registration) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	r.id = id
	return nil
}

// SetCreatedAt is used by repository when loading from database
func (r *Registration) SetCreatedAt(createdAt time.Time) {
	r.createdAt = createdAt
}
