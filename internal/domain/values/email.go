package values

import (
	"hinsun-backend/internal/core/failure"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

const (
	EmailMaxLength = 254
)

type Email struct {
	value string
}

func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, failure.NewValidationFailure("email cannot be empty")
	}

	if !emailRegex.MatchString(email) {
		return nil, failure.NewValidationFailure("invalid email format")
	}

	if len(email) > EmailMaxLength {
		return nil, failure.NewValidationFailure("email exceeds maximum length")
	}

	return &Email{value: email}, nil
}

func (e *Email) LocalPart() string {
	for i, char := range e.value {
		if char == '@' {
			return e.value[:i]
		}
	}

	return e.value // Fallback, should not happen if email is valid
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	return e.value == other.value
}
