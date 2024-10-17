package arraystring

import "testing"

var mergeStringsAlternatelyTestCases = []struct {
	Name  string
	Input []string
	Want  string
}{
	{"Example 1", []string{"abc", "pqr"}, "apbqcr"},
	{"Example 2", []string{"ab", "pqrs"}, "apbqrs"},
	{"Example 3", []string{"abcd", "pq"}, "apbqcd"},
}

func TestMergeStringsAlternately(t *testing.T) {
	for _, tc := range mergeStringsAlternatelyTestCases {
		got := mergeAlternately(tc.Input[0], tc.Input[1])
		if tc.Want != got {
			t.Error("want", tc.Want, "got", got)
		}
	}
}