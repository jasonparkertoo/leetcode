package stack

import "testing"

var decodeStringTestCases = []struct{
	Name string
	Have string
	Want string
}{
	{"1", "3[a]2[bc]", "aaabcbc"},
	{"1", "3[a2[c]]", "accaccacc"},
	{"1", "2[abc]3[cd]ef", "abcabccdcdcdef"},
}

func TestDecodeString(t *testing.T) {
	for _, tc := range decodeStringTestCases {
		got := decodeString(tc.Have)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}