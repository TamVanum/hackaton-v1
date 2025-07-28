package domain

import "errors"

type Role struct {
	id   int
	name string
}

func NewRole(id int, name string) *Role {
	return &Role{id: id, name: name}
}

func (r *Role) ID() int {
	return r.id
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) SetID(id int) error {
	if id <= 0 {
		return errors.New("id must be greater than 0")
	}
	r.id = id
	return nil
}

func (r *Role) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	r.name = name
	return nil
}
