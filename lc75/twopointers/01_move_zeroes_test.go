package twopointers

import "testing"

var moveZerosTestCases = []struct {
	Name  string
	Input []int
	Want  []int
}{
	{"1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{"2", []int{0}, []int{0}},
	{"3", []int{2}, []int{2}},
	{"4", []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
	{"5", []int{0, 0, 1}, []int{1, 0, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, tc := range moveZerosTestCases {
		moveZeroes(tc.Input)
		if len(tc.Want) != len(tc.Input) {
			t.Error("test", tc.Name, "want", tc.Want, "got", tc.Input)
		}
		for elem := range tc.Want {
			if tc.Want[elem] != tc.Input[elem] {
				t.Error("test", tc.Name, "want", tc.Want[elem], "got", tc.Input[elem])
			}
		}
	}
}