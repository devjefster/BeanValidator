package validator

import (
	"testing"
	"time"
)

// Test required rule
func TestRequiredRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{"", "TestField is required"},
		{"hello", ""},
		{nil, "TestField is required"},
	}

	for _, test := range tests {
		err := requiredRule("TestField", test.value)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test non-null rule
func TestNonNullRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{nil, "TestField must not be null"},
		{"hello", ""},
		{123, ""},
	}

	for _, test := range tests {
		err := nonNullRule("TestField", test.value)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test non-blank rule
func TestNonBlankRule(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"", "TestField must not be blank"},
		{"   ", "TestField must not be blank"},
		{"hello", ""},
	}

	for _, test := range tests {
		err := nonBlankRule("TestField", test.value)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test non-empty rule
func TestNonEmptyRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{[]string{}, "TestField must not be empty"},
		{"hello", ""},
		{[]int{1, 2, 3}, ""},
	}

	for _, test := range tests {
		err := nonEmptyRule("TestField", test.value)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test min rule
func TestMinRule(t *testing.T) {
	tests := []struct {
		value    int
		param    string
		expected string
	}{
		{5, "10", "TestField must be at least 10"},
		{15, "10", ""},
	}

	for _, test := range tests {
		err := minRule("TestField", test.value, test.param)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test max rule
func TestMaxRule(t *testing.T) {
	tests := []struct {
		value    int
		param    string
		expected string
	}{
		{15, "10", "TestField must be at most 10"},
		{5, "10", ""},
	}

	for _, test := range tests {
		err := maxRule("TestField", test.value, test.param)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test email rule
func TestEmailRule(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"test@example.com", ""},
		{"invalid-email", "TestField is not a valid email"},
		{"", "TestField is not a valid email"},
	}

	for _, test := range tests {
		err := emailRule("TestField", test.value)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test date rule with custom format
func TestDateRule(t *testing.T) {
	tests := []struct {
		value    string
		format   string
		expected string
	}{
		{"2024-02-01", "2006-01-02", ""},                                           // Valid date
		{"2024-02-30", "2006-01-02", "TestField must match the format 2006-01-02"}, // Invalid date
		{"not-a-date", "2006-01-02", "TestField must match the format 2006-01-02"},
	}

	for _, test := range tests {
		err := dateRule("TestField", test.value, test.format)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test date format rule
func TestDateFormatRule(t *testing.T) {
	tests := []struct {
		value    string
		format   string
		expected string
	}{
		{"2024-02-01", "2006-01-02", ""}, // Matches YYYY-MM-DD
		{"01/02/2024", "02/01/2006", ""}, // Matches MM/DD/YYYY
		{"01-02-2024", "02-01-2006", ""}, // Matches MM-DD-YYYY
		{"not-a-date", "2006-01-02", "TestField must match the format 2006-01-02"},
	}

	for _, test := range tests {
		err := dateFormatRule("TestField", test.value, test.format)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test after date rule
func TestAfterDateRule(t *testing.T) {
	tests := []struct {
		value    string
		refDate  string
		format   string
		expected string
	}{
		{"2024-02-01", "2024-01-01", "2006-01-02", ""},
		{"2023-12-31", "2024-01-01", "2006-01-02", "TestField must be after 2024-01-01"},
	}

	for _, test := range tests {
		err := afterDateRule("TestField", test.value, test.refDate, test.format)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test before date rule
func TestBeforeDateRule(t *testing.T) {
	tests := []struct {
		value    string
		refDate  string
		format   string
		expected string
	}{
		{"2023-12-31", "2024-01-01", "2006-01-02", ""},
		{"2024-02-01", "2024-01-01", "2006-01-02", "TestField must be before 2024-01-01"},
	}

	for _, test := range tests {
		err := beforeDateRule("TestField", test.value, test.refDate, test.format)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test between date rule
func TestBetweenDateRule(t *testing.T) {
	tests := []struct {
		value    string
		start    string
		end      string
		format   string
		expected string
	}{
		{"2024-06-01", "2024-01-01", "2024-12-31", "2006-01-02", ""},
		{"2023-12-31", "2024-01-01", "2024-12-31", "2006-01-02", "TestField must be between 2024-01-01 and 2024-12-31"},
	}

	for _, test := range tests {
		err := betweenDateRule("TestField", test.value, test.start, test.end, test.format)
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test past date rule
func TestPastDateRule(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	tests := []struct {
		value    string
		expected string
	}{
		{yesterday, ""},
		{tomorrow, "TestField must be in the past"},
		{today, "TestField must be in the past"},
	}

	for _, test := range tests {
		err := pastDateRule("TestField", test.value, "2006-01-02")
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test future date rule
func TestFutureDateRule(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	tests := []struct {
		value    string
		expected string
	}{
		{tomorrow, ""},
		{yesterday, "TestField must be in the future"},
		{today, "TestField must be in the future"},
	}

	for _, test := range tests {
		err := futureDateRule("TestField", test.value, "2006-01-02")
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test past-inclusive date rule
func TestPastInclusiveDateRule(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	tests := []struct {
		value    string
		expected string
	}{
		{today, ""},
		{tomorrow, "TestField must be in the past or today"},
	}

	for _, test := range tests {
		err := pastInclusiveDateRule("TestField", test.value, "2006-01-02")
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test future-inclusive date rule
func TestFutureInclusiveDateRule(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	tests := []struct {
		value    string
		expected string
	}{
		{today, ""},
		{yesterday, "TestField must be in the future or today"},
	}

	for _, test := range tests {
		err := futureInclusiveDateRule("TestField", test.value, "2006-01-02")
		if test.expected == "" && err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if test.expected != "" && (err == nil || err.Error() != test.expected) {
			t.Errorf("expected error '%s', got '%v'", test.expected, err)
		}
	}
}

// Test isTrue rule
func TestIsTrueRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{true, ""},
		{false, "TestField must be true"},
		{nil, "TestField must be true"},
	}

	for _, test := range tests {
		err := isTrueRule("TestField", test.value)
		validateError(t, err, test.expected)
	}
}

// Test positive rule
func TestPositiveRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{5, ""},
		{-3, "TestField must be positive"},
		{0, "TestField must be positive"},
	}

	for _, test := range tests {
		err := positiveRule("TestField", test.value)
		validateError(t, err, test.expected)
	}
}

// Test negative rule
func TestNegativeRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{-5, ""},
		{3, "TestField must be negative"},
		{0, "TestField must be negative"},
	}

	for _, test := range tests {
		err := negativeRule("TestField", test.value)
		validateError(t, err, test.expected)
	}
}

// Test size rule
func TestSizeRule(t *testing.T) {
	tests := []struct {
		value    interface{}
		size     string
		expected string
	}{
		{[]int{1, 2, 3}, "3", ""},
		{"hello", "5", ""},
		{"hello", "3", "TestField must have exactly 3 elements"},
	}

	for _, test := range tests {
		err := sizeRule("TestField", test.value, test.size)
		validateError(t, err, test.expected)
	}
}

// Helper function to validate expected errors
func validateError(t *testing.T, err error, expected string) {
	if expected == "" && err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if expected != "" && (err == nil || err.Error() != expected) {
		t.Errorf("expected error '%s', got '%v'", expected, err)
	}
}
