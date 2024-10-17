package slidingwindow

import "testing"

var maximumNumberOfVowelsInASubstringTestCases = []struct {
	Name  string
	Input string
	K     int
	Want  int
}{
	{"1", "abciiidef", 3, 3},
	{"2", "aeiou", 2, 2},
	{"3", "leetcode", 3, 2},
	{"4", "weallloveyou", 7, 4},
	{"5", "a", 1, 1},
}

func TestMaximumNumberOfVowelsInASubstring(t *testing.T) {
	for _, tc := range maximumNumberOfVowelsInASubstringTestCases {
		got := maxVowels(tc.Input, tc.K)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
