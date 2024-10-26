package stack

import "testing"

var removeStarsFromAstringTestCases = []struct {
	Name  string
	Input string
	Want  string
}{
	{"1", "leet**cod*e", "lecoe"},
	{"1", "erase*****", ""},
}

func TestRemoveStarsFromAstring(t *testing.T) {
	for _, tc := range removeStarsFromAstringTestCases {
		got := removeStars(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}