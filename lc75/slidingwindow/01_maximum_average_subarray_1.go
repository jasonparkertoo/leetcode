package slidingwindow

// findMaxAverage calculates the maximum average of a contiguous subarray of size k in the given nums array.
// It uses a sliding window approach to efficiently compute the average for each subarray.
//
// Parameters:
//   - nums: A slice of integers representing the input array.
//   - k: An integer representing the size of the subarray.
//
// Returns:
//   - float64: The maximum average of any contiguous subarray of size k.
func findMaxAverage(nums []int, k int) float64 {
	// Initialize variables for maximum average, sum, and pointers
	var maxAvg, sum float64
	var ptr, right int

	// Calculate the sum of the first k elements
	for right < k {
		sum += float64(nums[right])
		if right == (k - 1) {
			// Calculate the average of the first window
			maxAvg = sum / float64(k)
			// Remove the first element from the sum
			sum -= float64(nums[ptr])
		}
		right += 1
	}
	ptr += 1

	// Slide the window through the remaining elements
	for right < len(nums) {
		// Add the next element to the sum
		sum += float64(nums[right])
		// Update the maximum average if necessary
		maxAvg = max(maxAvg, sum/float64(k))
		// Remove the first element of the window from the sum
		sum -= float64(nums[ptr])
		// Move the window
		right++
		ptr++
	}

	// Return the maximum average found
	return maxAvg
}