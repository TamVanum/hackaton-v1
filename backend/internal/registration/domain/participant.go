package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

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

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if strings.TrimSpace(nickname) == "" {
		return nil, errors.New("nickname cannot be empty")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if !isValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	if strings.TrimSpace(region) == "" {
		return nil, errors.New("region cannot be empty")
	}

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

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

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

func (p *Participant) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	p.id = id
	return nil
}

func (p *Participant) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}
