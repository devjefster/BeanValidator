package validator

import (
	"fmt"
	"time"
)

// dateRule validates if a string is a valid date based on a user-provided format
func dateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("date rule requires a format parameter (e.g., '2006-01-02')")
	}
	format := params[0] // User-provided date format

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	_, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}
	return nil
}

// dateFormatRule checks if a date matches a custom format
func dateFormatRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("date-format rule requires a format parameter")
	}
	format := params[0]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}
	_, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}
	return nil
}

// afterDateRule ensures a date is after a specific date
func afterDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 2 {
		return fmt.Errorf("after rule requires a reference date and format (e.g., '2024-01-01,2006-01-02')")
	}
	refDateStr, format := params[0], params[1]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	refDate, err := time.Parse(format, refDateStr)
	if err != nil {
		return fmt.Errorf("invalid reference date for %s", fieldName)
	}

	if !parsedValue.After(refDate) {
		return fmt.Errorf("%s must be after %s", fieldName, refDateStr)
	}
	return nil
}

// beforeDateRule ensures a date is before a specific date
func beforeDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 2 {
		return fmt.Errorf("before rule requires a reference date and format (e.g., '2024-01-01,2006-01-02')")
	}
	refDateStr, format := params[0], params[1]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	refDate, err := time.Parse(format, refDateStr)
	if err != nil {
		return fmt.Errorf("invalid reference date for %s", fieldName)
	}

	if !parsedValue.Before(refDate) {
		return fmt.Errorf("%s must be before %s", fieldName, refDateStr)
	}
	return nil
}

// betweenDateRule ensures a date is within a range
func betweenDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 3 {
		return fmt.Errorf("between rule requires a start date, end date, and format (e.g., '2024-01-01,2024-12-31,2006-01-02')")
	}
	startDateStr, endDateStr, format := params[0], params[1], params[2]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	startDate, err := time.Parse(format, startDateStr)
	if err != nil {
		return fmt.Errorf("invalid start date for %s", fieldName)
	}

	endDate, err := time.Parse(format, endDateStr)
	if err != nil {
		return fmt.Errorf("invalid end date for %s", fieldName)
	}

	if parsedValue.Before(startDate) || parsedValue.After(endDate) {
		return fmt.Errorf("%s must be between %s and %s", fieldName, startDateStr, endDateStr)
	}
	return nil
}
func pastDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("past rule requires a format parameter (e.g., '2006-01-02')")
	}
	format := params[0]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	today := time.Now().Truncate(24 * time.Hour)

	if !parsedValue.Before(today) {
		return fmt.Errorf("%s must be in the past", fieldName)
	}
	return nil
}

func futureDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("future rule requires a format parameter (e.g., '2006-01-02')")
	}
	format := params[0]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	today := time.Now().Truncate(24 * time.Hour)

	if !parsedValue.After(today) {
		return fmt.Errorf("%s must be in the future", fieldName)
	}
	return nil
}

// pastInclusiveDateRule ensures a date is in the past or today
func pastInclusiveDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("past-inclusive rule requires a format parameter (e.g., '2006-01-02')")
	}
	format := params[0]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	if parsedValue.After(time.Now()) {
		return fmt.Errorf("%s must be in the past or today", fieldName)
	}
	return nil
}

// futureInclusiveDateRule ensures a date is in the future or today
func futureInclusiveDateRule(fieldName string, value interface{}, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("future-inclusive rule requires a format parameter (e.g., '2006-01-02')")
	}
	format := params[0]

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("%s must be a string representing a date", fieldName)
	}

	parsedValue, err := time.Parse(format, str)
	if err != nil {
		return fmt.Errorf("%s must match the format %s", fieldName, format)
	}

	if parsedValue.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("%s must be in the future or today", fieldName)
	}
	return nil
}
