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

func NewNotFoundFailure(message string) *Failure {
	return NewFailure(NotFoundFailure, message)
}

func NewValidationFailure(message string) *Failure {
	return NewFailure(ValidationFailure, message)
}

func NewDatabaseFailure(message string) *Failure {
	return NewFailure(DatabaseFailure, message)
}

func NewConflictFailure(message string) *Failure {
	return NewFailure(ConflictFailure, message)
}

func NewInternalFailure(message string, cause error) *Failure {
	return NewFailure(InternalFailure, message).WithCause(cause)
}

func NewAuthenticationFailure(message string) *Failure {
	return NewFailure(UnauthorizedFailure, message)
}
