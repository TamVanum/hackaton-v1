package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type Registration struct {
	id              int
	name            string
	nickname        string
	email           string
	region          string
	projectIdea     string
	teamPreference  bool
	desiredTeammate *string
	createdAt       time.Time
}

func NewRegistration(name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*Registration, error) {
	// Validate name
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	// Validate nickname
	if strings.TrimSpace(nickname) == "" {
		return nil, errors.New("nickname cannot be empty")
	}

	// Validate email
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	// Validate region
	if strings.TrimSpace(region) == "" {
		return nil, errors.New("region cannot be empty")
	}

	// Validate project idea
	if strings.TrimSpace(projectIdea) == "" {
		return nil, errors.New("project idea cannot be empty")
	}

	return &Registration{
		name:            strings.TrimSpace(name),
		nickname:        strings.TrimSpace(nickname),
		email:           email,
		region:          strings.TrimSpace(region),
		projectIdea:     strings.TrimSpace(projectIdea),
		teamPreference:  teamPreference,
		desiredTeammate: desiredTeammate,
		createdAt:       time.Now(),
	}, nil
}

// isValidEmail validates email format using regex
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
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

func (r *Registration) Email() string {
	return r.email
}

func (r *Registration) Region() string {
	return r.region
}

func (r *Registration) ProjectIdea() string {
	return r.projectIdea
}

func (r *Registration) TeamPreference() bool {
	return r.teamPreference
}

func (r *Registration) DesiredTeammate() *string {
	return r.desiredTeammate
}

func (r *Registration) CreatedAt() time.Time {
	return r.createdAt
}

func (r *Registration) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	r.id = id
	return nil
}

func (r *Registration) SetCreatedAt(createdAt time.Time) {
	r.createdAt = createdAt
}
