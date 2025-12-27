package https

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// BindQuery binds URL query parameters to the target struct
// Supports: string, []string (comma-separated), int, bool
// Use struct tag `query:"param_name"` to specify the query parameter name
func BindQuery(r *http.Request, target interface{}) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("target must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()
	query := r.URL.Query()

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get query parameter name from struct tag or use field name
		queryTag := fieldType.Tag.Get("query")
		if queryTag == "" {
			continue // Skip fields without query tag
		}

		queryValue := query.Get(queryTag)
		if queryValue == "" {
			continue // Skip empty values
		}

		// Set field value based on type
		if !field.CanSet() {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(queryValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(queryValue, 10, 64)
			if err != nil {
				return errors.New("invalid integer value for " + queryTag)
			}

			field.SetInt(intVal)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(queryValue, 10, 64)
			if err != nil {
				return errors.New("invalid unsigned integer value for " + queryTag)
			}

			field.SetUint(uintVal)

		case reflect.Bool:
			boolVal, err := strconv.ParseBool(queryValue)
			if err != nil {
				return errors.New("invalid boolean value for " + queryTag)
			}

			field.SetBool(boolVal)

		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.String {
				// Handle comma-separated string slice
				values := ParseCommaSeparated(queryValue)
				field.Set(reflect.ValueOf(values))
			}
		}
	}

	return nil
}

// ParseCommaSeparated parses a comma-separated string into a slice of strings
// Trims whitespace from each value
func ParseCommaSeparated(s string) []string {
	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
