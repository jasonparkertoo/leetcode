package scratch
// slidingWindow calculates the maximum sum of a contiguous subarray of size 'window' within the given 'data' slice.
// It returns the maximum sum found, or -1 if the data slice is smaller than the window size.
//
// Parameters:
//   - data: A slice of integers to analyze.
//   - window: The size of the sliding window.
//
// Returns:
//   The maximum sum of any contiguous subarray of size 'window', or -1 if the input is invalid.
func slidingWindow(data []int, window int) int {

	// check that data is greater than the window size
	if len(data) < window {
		return -1
	}

	// Initialize sum and maxSum. maxSum is set to the smallest possible int64 value
	sum, maxSum := 0, -1<<63

	// Calculate the sum of the first window
	for _, n := range data[:window] {
		sum += n
	}

	// Update maxSum if the first window sum is greater
	if sum > maxSum {
		maxSum = sum
	}

	// Slide the window through the rest of the data
	for i, n := range data[window:] {
		// Remove the first element of the previous window
		sum -= data[window - window + i]
		// Add the new element to the window
		sum += n
		// Update maxSum if the current window sum is greater
		if sum > maxSum {
			maxSum = sum
		}
	}

	// Return the maximum sum found
	return maxSum
}