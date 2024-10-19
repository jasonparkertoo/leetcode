package hashmapset

import (
	"testing"
)

var findDifferenceOfTwoArraysTestCases = []struct {
	Name  string
	Input [][]int
	Want  [][]int
}{
	{"1", [][]int{{1, 2, 3}, {2, 4, 6}}, [][]int{{1, 3}, {4, 6}}},
	{"1", [][]int{{1, 2, 3, 3}, {1, 1, 2, 2}}, [][]int{{3}, {}}},
}

func TestFindDifferenceOfTwoArray(t *testing.T) {
	for _, tc := range findDifferenceOfTwoArraysTestCases {
		got := findDifference(tc.Input[0], tc.Input[1])
		for x := range tc.Input {
			for y := range tc.Want[x] {
				if tc.Want[x][y] != got[x][y] {
					t.Error("test", tc.Name, "want", tc.Want[x][y], "got", got[x][y])
				}
			}
		}
	}
}
