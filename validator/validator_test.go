package validator

import (
	"testing"
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
