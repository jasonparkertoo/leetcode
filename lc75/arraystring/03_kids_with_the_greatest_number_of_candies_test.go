package arraystring

import "testing"

var kidsWithGreatestNumberOfCandiesTestCases = []struct {
	Name         string
	Candies      []int
	ExtraCandies int
	Want         []bool
}{
	{"1", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{"2", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{"3", []int{12, 1, 12}, 10, []bool{true, false, true}},
}

func TestKidsWithGreatestNumberOfCandies(t *testing.T) {

	for _, tc := range kidsWithGreatestNumberOfCandiesTestCases {
		got := kidsWithCandies(tc.Candies, tc.ExtraCandies)
		for i := range tc.Want {
			if len(tc.Want) != len(got) {
				t.Error("test", tc.Name, "want", tc.Want, "got", got)
			}
			if tc.Want[i] != got[i] {
				t.Error("test", tc.Name, "want", tc.Want, "got", got)
			}
		}
	}
}