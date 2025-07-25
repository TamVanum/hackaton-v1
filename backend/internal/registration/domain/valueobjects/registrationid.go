package valueobjects

import "errors"

type RegistrationID struct {
	value int
}

func NewRegistrationID(value int) (*RegistrationID, error) {
	if value < 0 {
		return nil, errors.New("registration ID must be positive")
	}

	return &RegistrationID{value: value}, nil
}

func (r *RegistrationID) Value() int {
	return r.value
}
