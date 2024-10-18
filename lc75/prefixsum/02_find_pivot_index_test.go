package prefixsum

import "testing"

// findPivotIndexTestCases is a slice of test cases for the pivotIndex function.
// Each test case consists of a name, input slice of integers, and the expected output.
var findPivotIndexTestCases = []struct {
	Name  string
	Input []int
	Want  int
}{
	{"1", []int{1, 7, 3, 6, 5, 6}, 3},
	{"1", []int{1, 2, 3}, -1},
	{"1", []int{2, 1, -1}, 0},
}

// TestFindPivotIndex tests the pivotIndex function using the predefined test cases.
// It iterates through each test case, calls the pivotIndex function with the input,
// and compares the result with the expected output. If there's a mismatch,
// it reports an error with the test case name, expected value, and actual value.
func TestFindPivotIndex(t *testing.T) {
	for _, tc := range findPivotIndexTestCases {
		got := pivotIndex(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}