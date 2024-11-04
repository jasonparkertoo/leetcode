package queue

import "testing"

var dota2SenateTestCases = []struct{
	Name string
	Have string
	Want string
}{
	{"1", "RD", "Radiant"},
}

func TestPredictPartyVictory(t *testing.T) {
	for _, tc := range dota2SenateTestCases {
		got := predictPartyVictory(tc.Have)
		if tc.Want != got {
			t.Error("test", tc.Name, "want", tc.Want, "have", tc.Have)
		}
	}
}