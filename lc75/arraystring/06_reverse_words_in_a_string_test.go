package arraystring

import "testing"

var reverseWordsInAStringTestCases = []struct {
	Name  string
	Input string
	Want  string
}{
	{"1", "the sky is blue", "blue is sky the"},
	{"2", "  hello world  ", "world hello"},
	{"3", "a good   example", "example good a"},
	{"4", "F R  I   E    N     D      S      ", "S D N E I R F"},
}

func TestReverseWordsInAString(t *testing.T) {
	for _, tc := range reverseWordsInAStringTestCases {
		got := reverseWords(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}
