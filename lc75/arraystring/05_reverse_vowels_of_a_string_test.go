package arraystring

import "testing"

var reverseVowelOfAStringTestCases = []struct {
	Name  string
	Input string
	Want  string
}{
	{"1", "IceCreAm", "AceCreIm"},
	{"2", "leetcode", "leotcede"},
	{"3", "lrrtcfde", "lrrtcfde"},
	{"4", "a a", "a a"},
}

func TestReverseVowelOfAString(t *testing.T) {
	for _, tc := range reverseVowelOfAStringTestCases {
		got := reverseVowels(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}