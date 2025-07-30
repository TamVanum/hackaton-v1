package domain

import (
	"errors"
	"strings"
)

// Role represents a participant's role in the hackathon
type Role struct {
	id          int
	name        string
	description string
}

func NewRole(name, description string) (*Role, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("role name cannot be empty")
	}

	if strings.TrimSpace(description) == "" {
		return nil, errors.New("role description cannot be empty")
	}

	return &Role{
		name:        strings.TrimSpace(name),
		description: strings.TrimSpace(description),
	}, nil
}

func (r *Role) ID() int {
	return r.id
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) Description() string {
	return r.description
}

func (r *Role) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	r.id = id
	return nil
}

func GetDefaultRoles() []*Role {
	return []*Role{
		{id: 1, name: "Frontend Developer", description: "Specializes in user interface and user experience development"},
		{id: 2, name: "Backend Developer", description: "Focuses on server-side logic, databases, and API development"},
		{id: 3, name: "Fullstack Developer", description: "Works on both frontend and backend development"},
		{id: 4, name: "Designer", description: "Focuses on UI/UX design and user experience"},
		{id: 5, name: "Product Manager", description: "Manages product development and strategy"},
		{id: 6, name: "Data Scientist", description: "Analyzes data and builds machine learning models"},
		{id: 7, name: "DevOps Engineer", description: "Manages infrastructure and deployment pipelines"},
		{id: 8, name: "Other", description: "Other technical or non-technical roles"},
	}
}
