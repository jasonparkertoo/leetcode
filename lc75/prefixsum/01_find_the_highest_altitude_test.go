package prefixsum

import "testing"

// largestGainTestCases is a slice of test cases for the largestAltitude function.
// Each test case contains a name, an input slice of integers, and the expected output.
var largestGainTestCases = []struct {
	Name  string
	Input []int
	Want  int
}{
	{"1", []int{-5, 1, 5, 0, -7}, 1},
	{"2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
}

// TestLargestGainTestCases runs through a series of test cases to verify the correctness
// of the largestAltitude function. It iterates over the predefined largestGainTestCases,
// applying the largestAltitude function to each input and comparing the result with the
// expected output. If any discrepancies are found, it reports an error using the testing
// framework, providing details about the failed test case.
func TestLargestGainTestCases(t *testing.T) {
	for _, tc := range largestGainTestCases {
		got := largestAltitude(tc.Input)
		if tc.Want != got {
			t.Error("test case", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
