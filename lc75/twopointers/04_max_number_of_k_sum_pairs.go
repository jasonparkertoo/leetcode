package twopointers

// maxOperations counts the maximum number of operations that can be performed on the given array
// to create pairs of numbers that sum up to k.
//
// Parameters:
//   - nums: A slice of integers representing the input array.
//   - k: An integer representing the target sum for each pair.
//
// Returns:
//
//	An integer representing the maximum number of operations that can be performed.
//
// The function uses a map to keep track of the count of each number in the array.
// It iterates through the array, checking if the complement (k - num) exists in the map.
// If it does, an operation is counted, and the complement's count is decremented.
// If not, the current number's count is incremented in the map.
func maxOperations(nums []int, k int) int {
	counts := make(map[int]int, len(nums))
	operations := 0

	for _, num := range nums {
		complement := k - num
		if counts[complement] > 0 {
			operations++
			counts[complement]--
		} else {
			counts[num]++
		}
	}

	return operations
}