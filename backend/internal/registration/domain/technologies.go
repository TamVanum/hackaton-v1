package domain

import "errors"

type Technologies struct {
	id   int
	name string
}

func NewTechnologies(id int, name string) *Technologies {
	return &Technologies{id: id, name: name}
}

func (t *Technologies) ID() int {
	return t.id
}

func (t *Technologies) Name() string {
	return t.name
}

func (t *Technologies) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}
	t.id = id
	return nil
}

func (t *Technologies) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	t.name = name
	return nil
}
