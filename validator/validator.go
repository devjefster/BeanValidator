package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// Validate validates the fields of a struct based on tags
func Validate(s interface{}) ValidationErrors {
	var errs ValidationErrors
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		panic("validator.Validate: input must be a struct or a pointer to a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		// Skip unexported fields
		if !value.CanInterface() {
			continue
		}

		// Parse "validate" tag
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}

		// Apply validation rules
		ruleList := strings.Split(tag, ",")
		for _, rule := range ruleList {
			ruleName, params := parseRule(rule)
			if ruleFunc, exists := ValidationRules[ruleName]; exists {
				if err := ruleFunc(field.Name, value.Interface(), params...); err != nil {
					errs = append(errs, ValidationError{
						Field:   field.Name,
						Message: err.Error(),
					})
				}
			} else {
				errs = append(errs, ValidationError{
					Field:   field.Name,
					Message: fmt.Sprintf("unknown validation rule: %s", ruleName),
				})
			}
		}
	}
	return errs
}

// parseRule splits a rule into its name and parameters
func parseRule(rule string) (string, []string) {
	parts := strings.SplitN(rule, "=", 2)
	if len(parts) == 2 {
		return parts[0], strings.Split(parts[1], ",")
	}
	return parts[0], nil
}
