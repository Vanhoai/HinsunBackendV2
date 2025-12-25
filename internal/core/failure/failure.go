package failure

import (
	"fmt"
)

// FailureCode defines the type of domain failure
type FailureCode string

const (
	// ValidationFailure represents validation errors in domain objects
	ValidationFailure FailureCode = "VALIDATION_FAILURE"

	// NotFoundFailure represents entity not found errors
	NotFoundFailure FailureCode = "NOT_FOUND"

	// AlreadyExistsFailure represents duplicate entity errors
	AlreadyExistsFailure FailureCode = "ALREADY_EXISTS"

	// UnauthorizedFailure represents authorization errors
	UnauthorizedFailure FailureCode = "UNAUTHORIZED"

	// ForbiddenFailure represents permission errors
	ForbiddenFailure FailureCode = "FORBIDDEN"

	// ConflictFailure represents business rule conflicts
	ConflictFailure FailureCode = "CONFLICT"

	// InvalidOperationFailure represents invalid state operations
	InvalidOperationFailure FailureCode = "INVALID_OPERATION"

	// DomainRuleViolationFailure represents business rule violations
	DomainRuleViolationFailure FailureCode = "DOMAIN_RULE_VIOLATION"

	// InternalFailure represents internal domain errors
	InternalFailure FailureCode = "INTERNAL_FAILURE"
)

type Failure struct {
	Code    FailureCode
	Message string
	Details map[string]interface{}
	Cause   error
}

// Error implements the error interface for Failure
func (f *Failure) Error() string {
	if f.Cause != nil {
		return fmt.Sprintf("[%s]: %s (caused by: %v)", f.Code, f.Message, f.Cause)
	}

	return fmt.Sprintf("[%s]: %s", f.Code, f.Message)
}

// WithDetails adds details to the failure
func (f *Failure) WithDetails(key string, value interface{}) *Failure {
	if f.Details == nil {
		f.Details = make(map[string]interface{})
	}

	f.Details[key] = value
	return f
}

// WithCause adds a cause error to the failure
func (f *Failure) WithCause(err error) *Failure {
	f.Cause = err
	return f
}

// Is check if the error is a Failure
func Is(err error, code FailureCode) bool {
	if failure, ok := err.(*Failure); ok {
		return failure.Code == code
	}

	return false
}

// AsFailure attempts to cast an error to a Failure
func AsFailure(err error) (*Failure, bool) {
	if err == nil {
		return nil, false
	}
	f, ok := err.(*Failure)
	return f, ok
}
