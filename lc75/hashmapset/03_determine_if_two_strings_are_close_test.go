package hashmapset

import "testing"

var determineIfTwoStringsAreCloseTestCases = []struct {
	Name  string
	Word1 string
	Word2 string
	Want  bool
}{
	{"1", "abc", "bca", true},
	{"2", "a", "aa", false},
	{"3", "cabbba", "abbccc", true},
}

func TestDetermineIfTwoStringsAreClose(t *testing.T) {
	for _, tc := range determineIfTwoStringsAreCloseTestCases {
		got := closeStrings(tc.Word1, tc.Word2)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
