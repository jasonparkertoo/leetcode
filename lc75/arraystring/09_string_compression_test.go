package arraystring

import (
	"testing"
)

var stringCompressTestCases = []struct {
	Name  string
	Input []byte
	Want  int
}{
	{"1", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
	{"2", []byte{'a'}, 1},
	{"3", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
}

func TestStringCompress(t *testing.T) {
	for _, tc := range stringCompressTestCases {
		got := compress(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
	}
}