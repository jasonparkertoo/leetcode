package twopointers

// moveZeroes moves all the zeros in the given array to the end while maintaining
// the relative order of non-zero elements.
//
// Parameters:
//   - nums: A slice of integers to be modified in-place.
//
// The function uses a two-pointer approach:
// - 'last' keeps track of the position where the next non-zero element should be placed.
// - 'i' iterates through the array.
//
// Non-zero elements are moved to their correct positions, and then the remaining
// positions are filled with zeros.
func isSubsequence(s string, t string) bool {
	left, right, found := 0, 0, 0
	if len(s) == 0 {
		return true
	}
	for {
		if right == len(t) {
			return false
		}
		if s[left] != t[right] {
			right++
			continue
		}
		found++
		if found == len(s) {
			return true
		}
		right++
		left++
	}
}