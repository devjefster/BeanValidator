package validator

import (
	"fmt"
	"strings"
)

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface for ValidationError
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface for ValidationErrors
func (errs ValidationErrors) Error() string {
	var messages []string
	for _, err := range errs {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

// HasErrors checks if there are any validation errors
func (errs ValidationErrors) HasErrors() bool {
	return len(errs) > 0
}
