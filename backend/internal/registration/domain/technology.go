package domain

import (
	"errors"
	"strings"
)

// Technology represents a technology that participants are familiar with
type Technology struct {
	id          int
	name        string
	description string
	category    string // e.g., "Frontend", "Backend", "Database", "DevOps", etc.
}

func NewTechnology(name, description, category string) (*Technology, error) {
	// Validate name
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("technology name cannot be empty")
	}

	// Validate description
	if strings.TrimSpace(description) == "" {
		return nil, errors.New("technology description cannot be empty")
	}

	// Validate category
	if strings.TrimSpace(category) == "" {
		return nil, errors.New("technology category cannot be empty")
	}

	return &Technology{
		name:        strings.TrimSpace(name),
		description: strings.TrimSpace(description),
		category:    strings.TrimSpace(category),
	}, nil
}

// Getters
func (t *Technology) ID() int {
	return t.id
}

func (t *Technology) Name() string {
	return t.name
}

func (t *Technology) Description() string {
	return t.description
}

func (t *Technology) Category() string {
	return t.category
}

// SetID is used by repository after persistence
func (t *Technology) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	t.id = id
	return nil
}

// GetDefaultTechnologies returns predefined technologies for the hackathon
func GetDefaultTechnologies() []*Technology {
	return []*Technology{
		// Frontend
		{id: 1, name: "React", description: "JavaScript library for building user interfaces", category: "Frontend"},
		{id: 2, name: "Vue.js", description: "Progressive JavaScript framework", category: "Frontend"},
		{id: 3, name: "Angular", description: "Platform for building mobile and desktop web applications", category: "Frontend"},
		{id: 4, name: "TypeScript", description: "Typed superset of JavaScript", category: "Frontend"},
		{id: 5, name: "HTML/CSS", description: "Markup and styling languages for web development", category: "Frontend"},
		{id: 6, name: "Svelte", description: "Cybernetically enhanced web apps", category: "Frontend"},

		// Backend
		{id: 7, name: "Node.js", description: "JavaScript runtime for server-side development", category: "Backend"},
		{id: 8, name: "Go", description: "Statically typed, compiled programming language", category: "Backend"},
		{id: 9, name: "Python", description: "High-level programming language", category: "Backend"},
		{id: 10, name: "Java", description: "Object-oriented programming language", category: "Backend"},
		{id: 11, name: "C#", description: "Modern, object-oriented programming language", category: "Backend"},
		{id: 12, name: "PHP", description: "Popular general-purpose scripting language", category: "Backend"},
		{id: 13, name: "Ruby", description: "Dynamic, open source programming language", category: "Backend"},

		// Database
		{id: 14, name: "PostgreSQL", description: "Open source relational database", category: "Database"},
		{id: 15, name: "MySQL", description: "Open source relational database management system", category: "Database"},
		{id: 16, name: "MongoDB", description: "Document-oriented NoSQL database", category: "Database"},
		{id: 17, name: "Redis", description: "In-memory data structure store", category: "Database"},
		{id: 18, name: "SQLite", description: "Lightweight relational database", category: "Database"},

		// DevOps & Cloud
		{id: 19, name: "Docker", description: "Platform for developing, shipping, and running applications", category: "DevOps"},
		{id: 20, name: "Kubernetes", description: "Container orchestration platform", category: "DevOps"},
		{id: 21, name: "AWS", description: "Amazon Web Services cloud platform", category: "Cloud"},
		{id: 22, name: "Google Cloud", description: "Google Cloud Platform", category: "Cloud"},
		{id: 23, name: "Azure", description: "Microsoft Azure cloud platform", category: "Cloud"},

		// Tools & Others
		{id: 24, name: "Git", description: "Distributed version control system", category: "Tools"},
		{id: 25, name: "GraphQL", description: "Query language for APIs", category: "API"},
		{id: 26, name: "REST APIs", description: "Representational State Transfer APIs", category: "API"},
		{id: 27, name: "Machine Learning", description: "AI and machine learning technologies", category: "AI/ML"},
		{id: 28, name: "Blockchain", description: "Distributed ledger technology", category: "Blockchain"},
	}
}
