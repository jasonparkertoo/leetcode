package twopointers

import "testing"

var maxAreaTestCases = []struct {
	Name  string
	Input []int
	Want  int
}{
	{"1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{"2", []int{1, 1}, 1},
}

func TestMaxArea(t *testing.T) {
	for _, tc := range maxAreaTestCases {
		got := maxArea(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
