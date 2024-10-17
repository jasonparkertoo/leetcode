package arraystring

import "testing"

var increaseTripletSubSequenceTestCases = []struct {
	Name  string
	Input []int
	Want  bool
}{
	{"1", []int{1, 2, 3, 4, 5}, true},
	{"2", []int{5, 4, 3, 2, 1}, false},
	{"3", []int{2, 1, 5, 0, 4, 6}, true},
	{"4", []int{1, 2, 1, 3}, true},
	{"5", []int{20, 100, 10, 12, 5, 13}, true},
	{"6", []int{0, 4, 2, 1, 0, -1, -3}, false},
}

func TestIncreaseTripletSubSequence(t *testing.T) {
	for _, tc := range increaseTripletSubSequenceTestCases {
		got := increasingTriplet(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}