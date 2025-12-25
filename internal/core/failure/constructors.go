package failure

// NewFailure creates a new Failure instance
func NewFailure(code FailureCode, message string) *Failure {
	return &Failure{
		Code:    code,
		Message: message,
		Details: nil,
		Cause:   nil,
	}
}

// Constructors for common failures
func NewNotFoundFailure(message string) *Failure {
	return NewFailure(NotFoundFailure, message)
}

func NewValidationFailure(message string) *Failure {
	return NewFailure(ValidationFailure, message)
}
