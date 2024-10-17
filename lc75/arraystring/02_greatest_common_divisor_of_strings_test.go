package arraystring

import "testing"

var greatestCommonDivisorOfStringstestCase = []struct {
	Name  string
	Input []string
	Want  string
}{
	{"1", []string{"ABCABC", "ABC"}, "ABC"},
	{"2", []string{"ABABAB", "ABAB"}, "AB"},
	{"3", []string{"LEET", "CODE"}, ""},
	{"4", []string{"ABCDEF", "ABC"}, ""},
	{"5", []string{"TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"}, "TAUXX"},
}

func TestGreatestCommonDivisorOfStrings(t *testing.T) {
	for _, tc := range greatestCommonDivisorOfStringstestCase {
		got := gcdOfStrings(tc.Input[0], tc.Input[1])
		if tc.Want != got {
			t.Error("tc", tc.Name, "- want", tc.Want, "- got", got)
		}
	}
}
