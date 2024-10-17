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
func moveZeroes(nums []int) {
	last := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[last] = nums[i]
			last++
		}
	}

	for last < len(nums) {
		nums[last] = 0
		last++
	}
}