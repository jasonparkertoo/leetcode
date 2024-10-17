package arraystring

import "testing"

var canPlaceFlowersTestCases = []struct {
	Name  string
	Input []int
	Num   int
	Want  bool
}{
	{"1", []int{1, 0, 0, 0, 1}, 1, true},
	{"2", []int{1, 0, 0, 0, 1}, 2, false},
	{"3", []int{0, 0, 1, 0, 0, 0, 1}, 2, true},
	{"4", []int{0, 0, 1, 0, 0, 0, 1, 0, 0}, 3, true},
	{"5", []int{1, 0, 0, 0, 0, 1}, 2, false},
}

func TestCanPlaceFlowers(t *testing.T) {
	for _, tc := range canPlaceFlowersTestCases {
		got := canPlaceFlowers(tc.Input, tc.Num)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}