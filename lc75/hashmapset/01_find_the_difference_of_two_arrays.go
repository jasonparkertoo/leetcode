package hashmapset

// findDifference takes two integer slices and returns a slice of two slices.
// The first slice contains elements that are in nums1 but not in nums2.
// The second slice contains elements that are in nums2 but not in nums1.
//
// Parameters:
//   - nums1: A slice of integers
//   - nums2: A slice of integers
//
// Returns:
//   - A slice of two integer slices representing the differences
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := toSet(nums1)
	set2 := toSet(nums2)

	answer := [][]int{{}, {}}
	answer[0] = diff(set1, set2)
	answer[1] = diff(set2, set1)

	return answer
}

// toSet converts a slice of integers to a set represented by a map.
//
// Parameters:
//   - arr: A slice of integers
//
// Returns:
//   - A map with integer keys and empty struct values, representing a set
func toSet(arr []int) map[int]struct{} {
	out := make(map[int]struct{})
	for _, e := range arr {
		out[e] = struct{}{}
	}
	return out
}

// diff returns a slice of integers that are in set 'a' but not in set 'b'.
//
// Parameters:
//   - a: A map representing the first set
//   - b: A map representing the second set
//
// Returns:
//   - A slice of integers representing the difference between sets 'a' and 'b'
func diff(a, b map[int]struct{}) []int {
	diff := make([]int, 0, len(a))
	for key := range a {
		if _, present := b[key]; !present {
			diff = append(diff, key)
		}
	}
	return diff
}
