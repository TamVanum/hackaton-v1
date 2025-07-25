package valueobjects

import "errors"

type Role struct {
	value string
}

const (
	RoleFrontend  = "frontend"
	RoleDesigner  = "designer"
	RoleBackend   = "backend"
	RoleFullstack = "fullstack"
)

func NewRole(value string) (*Role, error) {
	if !isValidRole(value) {
		return nil, errors.New("invalid role")
	}

	return &Role{value: value}, nil
}

func (r *Role) Value() string {
	return r.value
}

func (r *Role) String() string {
	return r.value
}

func isValidRole(role string) bool {
	switch role {
	case RoleFrontend, RoleDesigner, RoleBackend, RoleFullstack:
		return true
	default:
		return false
	}
}
