package hashmapset

// uniqueOccurrences determines if all numbers in the input array appear unique number of times
// For example: [1,2,2,1,1,3] has 1 appearing 3 times, 2 appearing 2 times, and 3 appearing 1 time
// Therefore uniqueOccurrences would return true since 1,2,3 times are all unique frequencies
// If [1,2,2,1,1,3,3] was passed, it would return false since 3 appears 2 times which is not unique
// 
// Parameters:
//   arr []int: Array of integers to check for unique occurrence counts
//
// Returns:
//   bool: true if all numbers appear unique number of times, false otherwise
func uniqueOccurrences(arr []int) bool {
	table := map[int]int{}
	for _, n := range arr {
		table[n] += 1
	}

	set := map[int]bool{}
	for _, count := range table {
		if set[count] {
			return false
		}
		set[count] = true
	}
	return true
}
