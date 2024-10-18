package prefixsum

import "testing"

var largestGainTestCases = []struct {
	Name  string
	Input []int
	Want  int
}{
	{"1", []int{-5, 1, 5, 0, -7}, 1},
	{"2", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
}

func TestLargestGainTestCases(t *testing.T) {
	for _, tc := range largestGainTestCases {
		got := largestAltitude(tc.Input)
		if tc.Want != got {
			t.Error("test case", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
