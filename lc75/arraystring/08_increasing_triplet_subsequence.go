package arraystring

// increasingTriplet checks if there exists an increasing subsequence of length 3 in the given array.
//
// Parameters:
//   - nums: A slice of integers to be checked.
//
// Returns:
//
//	A boolean value: true if an increasing triplet subsequence exists, false otherwise.
//
// The function uses two variables, 'first' and 'second', to keep track of the smallest
// and second smallest elements encountered so far. If a number greater than both 'first'
// and 'second' is found, it means an increasing triplet subsequence exists.
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first := (1 << 31) - 1
	second := first

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}