package arraystring

import "strconv"

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
func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	writeIndex := 0
	count := 1

	for i := 1; i <= len(chars); i++ {
		if i < len(chars) && chars[i] == chars[i-1] {
			count++
		} else {
			chars[writeIndex] = chars[i-1]
			writeIndex++

			if count > 1 {
				countStr := strconv.Itoa(count)
				for _, digit := range countStr {
					chars[writeIndex] = byte(digit)
					writeIndex++
				}
			}
			count = 1
		}
	}

	return writeIndex
}