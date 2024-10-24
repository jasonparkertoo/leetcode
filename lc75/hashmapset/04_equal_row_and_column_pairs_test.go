package hashmapset

import "testing"

var equalRowAndColumnPairsTestCases = []struct {
	Name  string
	Input [][]int
	Want  int
}{
	{"1", [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
	{"2", [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	{"3", [][]int{{2, 1}, {3, 32}}, 0},
}

func TestEqualRowAndColumnPairs(t *testing.T) {
	for _, tc := range equalRowAndColumnPairsTestCases {
		got := equalPairs(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}