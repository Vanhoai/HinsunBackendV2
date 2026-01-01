package https

import (
	"encoding/json"
	"hinsun-backend/internal/core/failure"
	"net/http"
	"reflect"
)

type ResponseCode = string

const (
	InternalServerErrorCode ResponseCode = "INTERNAL_SERVER_ERROR"
	ValidationFailureCode   ResponseCode = "VALIDATION_FAILURE"
	BadRequestCode          ResponseCode = "BAD_REQUEST"
)

func mapSuccessCode(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "OK"
	case http.StatusCreated:
		return "CREATED"
	case http.StatusAccepted:
		return "ACCEPTED"
	default:
		return "SUCCESS"
	}
}

func normalizeDataResponse(data any) any {
	if data == nil {
		return nil
	}

	val := reflect.ValueOf(data)
	// Check if it's a nil slice
	if val.Kind() == reflect.Slice && val.IsNil() {
		// Return empty slice of the same type
		return reflect.MakeSlice(val.Type(), 0, 0).Interface()
	}

	// Check if it's a pointer to a nil slice
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		elem := val.Elem()
		if elem.Kind() == reflect.Slice && elem.IsNil() {
			emptySlice := reflect.MakeSlice(elem.Type(), 0, 0)
			return emptySlice.Interface()
		}
	}

	return data
}

// Success sends a successful response with data
func ResponseSuccess(w http.ResponseWriter, statusCode int, message string, data any, meta ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// collect meta if provided
	var metaData any
	if len(meta) > 0 {
		metaData = meta[0]
	}

	response := Response{
		Code:    mapSuccessCode(statusCode),
		Message: message,
		Payload: normalizeDataResponse(data),
		Meta:    metaData,
	}

	json.NewEncoder(w).Encode(response)
}

// RespondWithFailure handles domain failures and converts them to HTTP responses
func RespondWithFailure(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	// Check if it's a domain failure
	if domainFailure, ok := failure.AsFailure(err); ok {
		statusCode, response := mapFailureToResponse(domainFailure)
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generic error handling for non-domain errors
	w.WriteHeader(http.StatusInternalServerError)
	response := FailureResponse{
		Code:    InternalServerErrorCode,
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := FailureResponse{
		Code:    BadRequestCode,
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}

func ValidationFailed(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := FailureResponse{
		Code:    ValidationFailureCode,
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}

func mapFailureCodeToHTTPStatus(code failure.FailureCode) int {
	switch code {
	case failure.ValidationFailure:
		return http.StatusBadRequest
	case failure.NotFoundFailure:
		return http.StatusNotFound
	case failure.AlreadyExistsFailure:
		return http.StatusConflict
	case failure.UnauthorizedFailure:
		return http.StatusUnauthorized
	case failure.ForbiddenFailure:
		return http.StatusForbidden
	case failure.ConflictFailure:
		return http.StatusConflict
	case failure.InvalidOperationFailure:
		return http.StatusBadRequest
	case failure.DomainRuleViolationFailure:
		return http.StatusUnprocessableEntity
	case failure.DatabaseFailure:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func mapFailureToResponse(f *failure.Failure) (int, FailureResponse) {
	statusCode := mapFailureCodeToHTTPStatus(f.Code)
	if f.Cause == nil {
		response := FailureResponse{
			Code:    string(f.Code),
			Message: f.Message,
		}

		return statusCode, response
	}

	response := FailureResponse{
		Code:     string(f.Code),
		Message:  f.Message,
		CausedBy: f.Cause.Error(),
	}

	return statusCode, response
}
