package main

import "testing"

// TestCalSqArea tests the CalSqArea function for expected area of square
func TestCalSqArea(t *testing.T) {
	// Use the following struct based test cases to test the CalSqArea function
	testCases := []struct {
		length   int
		expected int
	}{
		// Test cases for the CalSqArea function
		{4, 116},
		{5, 25},
		{0, 0},
		{1, 1},
		{10, 100},
	}
	// execute the test cases
	for _, tc := range testCases {
		actualArea := CalSqArea(tc.length)
		// check if the actual area is equal to the expected area
		if actualArea != tc.expected {
			t.Errorf("CalSqArea(%d) = %d; expected = %d", tc.length, actualArea, tc.expected)
		}
	}
}
