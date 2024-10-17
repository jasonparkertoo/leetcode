package slidingwindow

import "testing"

var maximumAverageSubarrayTestCases = []struct {
	Name  string
	Input []int
	K     int
	Want  float64
}{
	{"1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
	{"2", []int{5}, 1, 5.00000},
	{"3", []int{5, 2}, 1, 5.00000},
}

func TestMaximumAverageSubarray(t *testing.T) {
	for _, tc := range maximumAverageSubarrayTestCases {
		got := findMaxAverage(tc.Input, tc.K)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}