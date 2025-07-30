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
}

func NewTechnology(name, description string) (*Technology, error) {
	// Validate name
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("technology name cannot be empty")
	}

	// Validate description
	if strings.TrimSpace(description) == "" {
		return nil, errors.New("technology description cannot be empty")
	}

	return &Technology{
		name:        strings.TrimSpace(name),
		description: strings.TrimSpace(description),
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

func (t *Technology) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	t.id = id
	return nil
}

func GetDefaultTechnologies() []*Technology {
	return []*Technology{
		// Frontend
		{id: 1, name: "React", description: "JavaScript library for building user interfaces"},
		{id: 2, name: "Vue.js", description: "Progressive JavaScript framework"},
		{id: 3, name: "Angular", description: "Platform for building mobile and desktop web applications"},
		{id: 4, name: "TypeScript", description: "Typed superset of JavaScript"},
		{id: 5, name: "HTML/CSS", description: "Markup and styling languages for web development"},
		{id: 6, name: "Svelte", description: "Cybernetically enhanced web apps"},

		// Backend
		{id: 7, name: "Node.js", description: "JavaScript runtime for server-side development"},
		{id: 8, name: "Go", description: "Statically typed, compiled programming language"},
		{id: 9, name: "Python", description: "High-level programming language"},
		{id: 10, name: "Java", description: "Object-oriented programming language"},
		{id: 11, name: "C#", description: "Modern, object-oriented programming language"},
		{id: 12, name: "PHP", description: "Popular general-purpose scripting language"},
		{id: 13, name: "Ruby", description: "Dynamic, open source programming language"},

		// Database
		{id: 14, name: "PostgreSQL", description: "Open source relational database"},
		{id: 15, name: "MySQL", description: "Open source relational database management system"},
		{id: 16, name: "MongoDB", description: "Document-oriented NoSQL database"},
		{id: 17, name: "Redis", description: "In-memory data structure store"},
		{id: 18, name: "SQLite", description: "Lightweight relational database"},

		// DevOps & Cloud
		{id: 19, name: "Docker", description: "Platform for developing, shipping, and running applications"},
		{id: 20, name: "Kubernetes", description: "Container orchestration platform"},
		{id: 21, name: "AWS", description: "Amazon Web Services cloud platform"},
		{id: 22, name: "Google Cloud", description: "Google Cloud Platform"},
		{id: 23, name: "Azure", description: "Microsoft Azure cloud platform"},

		// Tools & Others
		{id: 24, name: "Git", description: "Distributed version control system"},
		{id: 25, name: "GraphQL", description: "Query language for APIs"},
		{id: 26, name: "REST APIs", description: "Representational State Transfer APIs"},
		{id: 27, name: "Machine Learning", description: "AI and machine learning technologies"},
		{id: 28, name: "Blockchain", description: "Distributed ledger technology"},
	}
}
