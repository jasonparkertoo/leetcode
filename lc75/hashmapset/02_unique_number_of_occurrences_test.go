package hashmapset

import "testing"

var uniqueNumberOfPccurrencesTestCases = []struct {
	Name  string
	Input []int
	Want  bool
}{
	{"1", []int{1, 2, 2, 1, 1, 3}, true},
	{"2", []int{1, 2}, false},
	{"3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
}

func TestXxx(t *testing.T) {
	for _, tc := range uniqueNumberOfPccurrencesTestCases {
		got := uniqueOccurrences(tc.Input)
		if tc.Want != got {
			t.Error("test", tc.Name, "wanted", tc.Want, "got", got)
		}
	}
}