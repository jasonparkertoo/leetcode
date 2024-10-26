package stack

import "testing"

var astroidCollisionTestCases = []struct {
	Name  string
	Input []int
	Want  []int
}{
	{"1", []int{5, 10, -5}, []int{5, 10}},
	{"2", []int{8, -8}, []int{}},
	{"3", []int{10, 2, -5}, []int{10}},
}

func TestAstroidCollision(t *testing.T) {
	for _, tc := range astroidCollisionTestCases {
		got := astroidCollision(tc.Input)
		for i := range tc.Want {
			if tc.Want[i] != got[i] {
				t.Error("test", tc.Name, "want", tc.Want, "got", got)
			} 
		}
	}
}