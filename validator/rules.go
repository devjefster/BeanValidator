package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Rule defines a validation function
type Rule func(fieldName string, value interface{}, params ...string) error

// Predefined validation rules
var ValidationRules = map[string]Rule{
	"required":  requiredRule,
	"non-null":  nonNullRule,
	"non-blank": nonBlankRule,
	"non-empty": nonEmptyRule,
	"min":       minRule,
	"max":       maxRule,
	"email":     emailRule,
}

// requiredRule checks if a value is not empty
func requiredRule(fieldName string, value interface{}, _ ...string) error {
	if isEmpty(value) {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

// nonNullRule checks if a value is not nil
func nonNullRule(fieldName string, value interface{}, _ ...string) error {
	if isNil(value) {
		return fmt.Errorf("%s must not be null", fieldName)
	}
	return nil
}

// nonBlankRule checks if a string value is not empty or whitespace-only
func nonBlankRule(fieldName string, value interface{}, _ ...string) error {
	if v, ok := value.(string); ok && strings.TrimSpace(v) == "" {
		return fmt.Errorf("%s must not be blank", fieldName)
	}
	return nil
}

// nonEmptyRule checks if a slice or map is not empty
func nonEmptyRule(fieldName string, value interface{}, _ ...string) error {
	if isEmpty(value) {
		return fmt.Errorf("%s must not be empty", fieldName)
	}
	return nil
}

// minRule checks if an integer value is at least a minimum
func minRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("min rule requires a parameter")
	}
	minimum, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("invalid min parameter for %s", fieldName)
	}
	if v, ok := value.(int); ok && v < minimum {
		return fmt.Errorf("%s must be at least %d", fieldName, minimum)
	}
	return nil
}

// maxRule checks if an integer value is at most a maximum
func maxRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("max rule requires a parameter")
	}
	maximum, err := strconv.Atoi(params[0])
	if err != nil {
		return fmt.Errorf("invalid max parameter for %s", fieldName)
	}
	if v, ok := value.(int); ok && v > maximum {
		return fmt.Errorf("%s must be at most %d", fieldName, maximum)
	}
	return nil
}

// emailRule checks if a string value is a valid email
func emailRule(fieldName string, value interface{}, _ ...string) error {
	if v, ok := value.(string); ok && !isValidEmail(v) {
		return fmt.Errorf("%s is not a valid email", fieldName)
	}
	return nil
}

// Helper functions
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		return strings.TrimSpace(val.String()) == ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return val.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	default:
		zero := reflect.Zero(val.Type()).Interface()
		return reflect.DeepEqual(value, zero)
	}
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}
	val := reflect.ValueOf(value)
	return val.Kind() == reflect.Ptr && val.IsNil()
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
