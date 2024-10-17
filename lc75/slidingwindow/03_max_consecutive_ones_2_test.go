package slidingwindow

import "testing"

var maxConsecutiveOnes3TestCases = []struct {
	Name  string
	Input []int
	K     int
	Want  int
}{
	{"1", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
	{"1", []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
}

func TestMaxConsecutiveOnes3(t *testing.T) {
	for _, tc := range maxConsecutiveOnes3TestCases {
		got := longestOnes(tc.Input, tc.K)
		if tc.Want != got {
			t.Error("want", tc.Want, "got", got)
		}
	}
}