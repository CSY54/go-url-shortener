package url

import (
	"testing"
)

func TestB64ToUint32(t *testing.T) {
	testcases := []struct{
		input string
		expected uint32
		hasError bool
	}{
		{"!@#", 0, true},
		{"AAAAAAE", 0, true},
		{"AQ", 1, false},
		{"_____w", 4294967295, false},
	}

	for _, tc := range testcases {
		if output, err := B64ToUint32(tc.input); output != tc.expected || (err != nil) != tc.hasError {
			t.Errorf("Expected: %q, %t, got: %q, %t", tc.expected, tc.hasError, output, (err != nil))
		}
	}
}

func TestUint32ToB64(t *testing.T) {
	testcases := []struct{
		input uint32
		expected string
		hasError bool
	}{
		{1, "AQ", false},
		{4294967295, "_____w", false},
	}

	for _, tc := range testcases {
		if output, err := Uint32ToB64(tc.input); output != tc.expected || (err != nil) != tc.hasError {
			t.Errorf("Expected: %q, %t, got: %q, %t", tc.expected, tc.hasError, output, (err != nil))
		}
	}
}
