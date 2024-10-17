package slidingwindow

import "testing"

// Package slidingwindow provides functions for solving sliding window problems.
//
// This package includes a test file for the longestSubarray function, which finds
// the length of the longest subarray of 1s after deleting one element from the input array.

var longestSubarrayOf1sAfterDeletingOneElementTestCases = []struct {
	Name  string
	Input []int
	Want  int
}{
	{"1", []int{1, 1, 0, 1}, 3},
	{"1", []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
	{"1", []int{1, 1, 1}, 2},
}

// TestLongestSubarrayOf1sAfterDeletingOneElement tests the longestSubarray function
// against a set of predefined test cases. It ensures that the function correctly
// calculates the length of the longest subarray of 1s after deleting one element
// from the input array. The test iterates through multiple scenarios, comparing
// the expected output with the actual result from the function.
func TestLongestSubarrayOf1sAfterDeletingOneElement(t *testing.T) {
	for _, tc := range longestSubarrayOf1sAfterDeletingOneElementTestCases {
		got := longestSubarray(tc.Input)
		if tc.Want != got {
			t.Error("want", tc.Want, "got", got)
		}
	}
}