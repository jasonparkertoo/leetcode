package twopointers

import "testing"

var isSubsequenceTestCases = []struct {
	Name  string
	Input []string
	Want  bool
}{
	{"1", []string{"abc", "ahbgdc"}, true},
	{"2", []string{"axc", "ahbgdc"}, false},
	{"3", []string{"", ""}, true},
	{"4", []string{"a", "a"}, true},
	{"5", []string{"", "ahbgdc"}, true},
	{"6", []string{"ahbgdc", ""}, false},
}

func TestIsSubsequence(t *testing.T) {
	for _, tc := range isSubsequenceTestCases {
		got := isSubsequence(tc.Input[0], tc.Input[1])
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}