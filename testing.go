package utils

import "testing"

// Confirm that a value is nil
func AssertNil(t *testing.T, val interface{}) {
	if val != nil {
		t.Fatalf("The specified value is not nil: %v", val)
	}
}

// Confirm that a value is not nil
func AssertNotNil(t *testing.T, val interface{}) {
	if val == nil {
		t.Fatal("The specified value is nil")
	}
}

// Confirm that a string var is empty
func AssertEmptyStr(t *testing.T, val string) {
	if val != "" {
		t.Fatalf("The specified string is not empty: %s", val)
	}
}

// Configm that a string var is not empty
func AssertNotEmptyStr(t *testing.T, val string) {
	if val == "" {
		t.Fatal("The specified string is empty")
	}
}
