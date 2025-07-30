package domain

import (
	"errors"
	"time"
)

// Registration represents the complete registration process as an aggregate root
// It ensures all business invariants are met before the registration can be considered valid
type Registration struct {
	id           int
	participant  *Participant
	roles        []*Role
	technologies []*Technology
	createdAt    time.Time
	completedAt  *time.Time
}

func NewRegistration(
	participant *Participant,
	roles []*Role,
	technologies []*Technology,
) (*Registration, error) {

	if err := validateRegistrationRules(roles, technologies); err != nil {
		return nil, err
	}

	return &Registration{
		participant:  participant,
		roles:        roles,
		technologies: technologies,
		createdAt:    time.Now(),
	}, nil
}

func validateRegistrationRules(roles []*Role, technologies []*Technology) error {
	if len(roles) == 0 {
		return errors.New("registration must have at least one role")
	}

	if len(technologies) == 0 {
		return errors.New("registration must have at least one technology")
	}

	for _, role := range roles {
		if role == nil {
			return errors.New("invalid role in registration")
		}
	}

	for _, tech := range technologies {
		if tech == nil {
			return errors.New("invalid technology in registration")
		}
	}

	return nil
}

func (r *Registration) ID() int {
	return r.id
}

func (r *Registration) Participant() *Participant {
	return r.participant
}

func (r *Registration) Roles() []*Role {
	rolesCopy := make([]*Role, len(r.roles))
	copy(rolesCopy, r.roles)
	return rolesCopy
}

func (r *Registration) Technologies() []*Technology {
	techCopy := make([]*Technology, len(r.technologies))
	copy(techCopy, r.technologies)
	return techCopy
}

func (r *Registration) CompletedAt() *time.Time {
	return r.completedAt
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

func (r *Registration) SetCompletedAt(completedAt *time.Time) {
	r.completedAt = completedAt
}
