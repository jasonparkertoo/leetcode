package slidingwindow

// longestSubarray finds the length of the longest subarray containing only 1s after deleting one element.
//
// Parameters:
//   - nums: A binary array (containing only 0s and 1s)
//
// Returns:
//
//	The length of the longest subarray of 1s after deleting one element from the input array.
//
// The function uses a sliding window approach to find the longest subarray with at most one zero.
// It then returns the length of this subarray minus one (to account for the deletion).
func longestSubarray(nums []int) int {

	// left: left pointer of the sliding window
	// right: right pointer of the sliding window
	// zeros: count of zeros in the current window
	// maxLen: length of the longest valid subarray found so far
	left, right, zeros, maxLen := 0, 0, 0, 0

	// Iterate through the array using the right pointer
	for right < len(nums) {
		// If we encounter a zero, increment the zero counter
		if nums[right] == 0 {
			zeros++
		}

		// If we have more than one zero in our current window,
		// shrink the window from the left until we have at most one zero
		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		// Update the maximum length of the subarray
		// We use right-left instead of right-left+1 because we need to delete one element
		maxLen = max(maxLen, right-left)
		right++
	}

	// Return the minimum of maxLen and len(nums)-1
	// This handles the case where the entire array is ones
	return min(maxLen, len(nums)-1)

}
