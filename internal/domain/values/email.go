package values

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	return e.value == other.value
}
