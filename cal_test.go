package main

import "testing"

// TestCalSqArea tests the CalSqArea function for correctness using table-driven tests.
func TestCalSqArea(t *testing.T) {
	testCases := []struct {
		side     int
		expected int
	}{
		{4, 16},
		{5, 25},
		{0, 0},
		{1, 1},
		{10, 100},
	}

	for _, tc := range testCases {
		actualArea := CalSqArea(tc.side)
		if actualArea != tc.expected {
			t.Errorf("CalSqArea(%d) = %d; want %d", tc.side, actualArea, tc.expected)
		}
	}
}
