package slidingwindow

// longestOnes finds the length of the longest contiguous subarray with at most 'k' zeros
// that can be flipped to ones.
//
// Parameters:
//   - nums: A binary array (containing only 0s and 1s)
//   - k: The maximum number of zeros that can be flipped to ones
//
// Returns:
//
//	The length of the longest contiguous subarray that contains only 1s after flipping at most k 0s
func longestOnes(nums []int, k int) int {
	// left: left pointer of the sliding window
	// right: right pointer of the sliding window
	// zeros: count of zeros in the current window
	// maxOnes: length of the longest valid subarray found so far
	left, right, zeros, maxOnes := 0, 0, 0, 0

	// Iterate through the array using the right pointer
	for right < len(nums) {
		// If the current element is 0, increment the zero count
		if nums[right] == 0 {
			zeros++
		}

		// If we have more zeros than allowed (k), shrink the window from the left
		for zeros > k {
			// If the element at the left pointer is 0, decrement the zero count
			if nums[left] == 0 {
				zeros--
			}
			// Move the left pointer to the right
			left++
		}

		// Update maxOnes if the current window is longer
		maxOnes = max(maxOnes, right-left+1)
		// Move the right pointer to expand the window
		right++
	}

	// Return the length of the longest valid subarray
	return maxOnes
}