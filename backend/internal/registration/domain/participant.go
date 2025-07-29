package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Participant is the aggregate root representing a hackathon participant
type Participant struct {
	id              int
	name            string
	nickname        string
	email           string
	region          string
	projectIdea     string
	teamPreference  bool // true = wants to work in teams, false = wants to work alone
	desiredTeammate *string
	createdAt       time.Time
}

func NewParticipant(name, nickname, email, region, projectIdea string, teamPreference bool, desiredTeammate *string) (*Participant, error) {
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

	return &Participant{
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
func (p *Participant) ID() int {
	return p.id
}

func (p *Participant) Name() string {
	return p.name
}

func (p *Participant) Nickname() string {
	return p.nickname
}

func (p *Participant) Email() string {
	return p.email
}

func (p *Participant) Region() string {
	return p.region
}

func (p *Participant) ProjectIdea() string {
	return p.projectIdea
}

func (p *Participant) TeamPreference() bool {
	return p.teamPreference
}

func (p *Participant) DesiredTeammate() *string {
	return p.desiredTeammate
}

func (p *Participant) CreatedAt() time.Time {
	return p.createdAt
}

// SetID is used by repository after persistence
func (p *Participant) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	p.id = id
	return nil
}

// SetCreatedAt is used by repository when loading from database
func (p *Participant) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}
