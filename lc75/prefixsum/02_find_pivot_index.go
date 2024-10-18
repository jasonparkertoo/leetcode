package prefixsum

// pivotIndex finds the pivot index in the given slice of integers.
// The pivot index is the index where the sum of all numbers to its left
// is equal to the sum of all numbers to its right.
// If no such index exists, it returns -1.
//
// Parameters:
//   - nums: A slice of integers to search for the pivot index.
//
// Returns:
//   - The pivot index if found, otherwise -1.
func pivotIndex(nums []int) int {
	// Calculate the total sum of all numbers in the slice
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// Initialize the left sum to 0
	leftSum := 0
	
	// Iterate through the slice
	for i, num := range nums {
		// Check if the current index is the pivot
		// by comparing left sum with right sum
		if leftSum == totalSum-leftSum-num {
			return i // Return the pivot index if found
		}
		leftSum += num // Update the left sum
	}

	// If no pivot index is found, return -1
	return -1
}
