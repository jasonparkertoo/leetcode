package scratch

import "testing"

var slidingWindowTestCases = []struct {
	Name  string
	Input []int
	K int
	Want  int
}{
	// positive test case
	{"1", []int{3,1,4,3,2,6,6,3,4,4}, 3, 15},
	// window larger than input length
	{"2", []int{3}, 3, -1},
	// single element input
	{"3", []int{3}, 1, 3},
	// empty input
	{"4", []int{}, 1, -1},
	// empty input
	{"5", []int{ -1, -2, -5, -9, 0, 6, -1}, 4, -4},
}

func TestSlidingWindow(t *testing.T) {
	for _, tc := range slidingWindowTestCases {
		got := slidingWindow(tc.Input, tc.K)
		if got != tc.Want {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}

