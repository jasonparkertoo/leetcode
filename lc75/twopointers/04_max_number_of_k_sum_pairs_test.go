package twopointers

import "testing"

var maxNumberOfKSumPairsTestCases = []struct {
	Name  string
	Input []int
	K     int
	Want  int
}{
	{"1", []int{1, 2, 3, 4}, 5, 2},
	{"2", []int{3, 1, 3, 4, 3}, 6, 1},
}

func TestMaxNumberOfKSumPairs(t *testing.T) {
	for _, tc := range maxNumberOfKSumPairsTestCases {
		got := maxOperations(tc.Input, tc.K)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}