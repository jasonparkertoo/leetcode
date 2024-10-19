package hashmapset

import (
	"slices"
	"testing"
)

// findDifferenceOfTwoArraysTestCases is a slice of test cases for the findDifference function.
// Each test case contains a name, input arrays, and expected output arrays.
var findDifferenceOfTwoArraysTestCases = []struct {
	Name  string
	Input [][]int
	Want  [][]int
}{
	//{"1", [][]int{{1, 2, 3}, {2, 4, 6}}, [][]int{{1, 3}, {4, 6}}},
	//{"1", [][]int{{1, 2, 3, 3}, {1, 1, 2, 2}}, [][]int{{3}, {}}},
	{"1", [][]int{{-80, -15, -81, -28, -61, 63, 14, -45, -35, -10}, {-1, -40, -44, 41, 10, -43, 69, 10, 2}}, [][]int{{-81, -35, -10, -28, -61, -45, -15, 14, -80, 63}, {-1, 2, 69, -40, 41, 10, -43, -44}}},
}

// findDifference finds the unique elements in two arrays.
// It returns a slice of two slices:
// - The first slice contains elements that are in nums1 but not in nums2.
// - The second slice contains elements that are in nums2 but not in nums1.
// Each returned slice contains unique elements in sorted order.
//
// Parameters:
//   - nums1: The first input array of integers.
//   - nums2: The second input array of integers.
//
// Returns:
//   - A slice of two slices containing the unique elements as described above.
func TestFindDifferenceOfTwoArray(t *testing.T) {
	for _, tc := range findDifferenceOfTwoArraysTestCases {
		got := findDifference(tc.Input[0], tc.Input[1])
		for x := range tc.Input {
			for _, y := range tc.Input[x] {
				if !slices.Contains(got[x], y) {
					t.Error("test", tc.Name, "want", tc.Want, "got", got)
				}
			}
		}
	}
}
