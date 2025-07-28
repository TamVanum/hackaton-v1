package domain

import "errors"

type Technology struct {
	id   int
	name string
}

func NewTechnology(id int, name string) *Technology {
	return &Technology{id: id, name: name}
}

func (t *Technology) ID() int {
	return t.id
}

func (t *Technology) Name() string {
	return t.name
}

func (t *Technology) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}
	t.id = id
	return nil
}

func (t *Technology) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	t.name = name
	return nil
}
