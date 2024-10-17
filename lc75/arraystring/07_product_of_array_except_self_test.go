package arraystring

import "testing"

var productOfArrayExceptSelfTestCases = []struct {
	Name  string
	Input []int
	Want  []int
}{
	{"1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{"1", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductOfArrayExceptSelf(t *testing.T) {
	for _, tc := range productOfArrayExceptSelfTestCases {
		got := productExceptSelf(tc.Input)
		if len(tc.Want) != len(got) {
			t.Error("test", tc.Name, "want", tc.Want, "got", got)
		}
		for i := range tc.Want {
			if tc.Want[i] != got[i] {
				t.Error("test", tc.Name, "want", tc.Want[i], "got", got[i])
			}
		}
	}
}