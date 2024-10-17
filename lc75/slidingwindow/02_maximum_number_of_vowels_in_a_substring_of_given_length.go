package slidingwindow

// findMaxAverage calculates the maximum average of a contiguous subarray of length k in the given array.
//
// Parameters:
//   - nums: A slice of integers representing the input array.
//   - k: An integer representing the length of the subarray.
//
// Returns:
//
//	A float64 value representing the maximum average of any contiguous subarray of length k.
//
// The function uses a sliding window approach to efficiently calculate the average of each subarray
// and keeps track of the maximum average encountered.
func maxVowels(s string, k int) int {
	// isVowel is a helper function that checks if a given byte is a vowel
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	// Initialize variables:
	// left: left pointer of the sliding window
	// right: right pointer of the sliding window
	// count: current count of vowels in the window
	// maxVowels: maximum count of vowels found so far
	left, right, count, maxVowels := 0, 0, 0, 0

	// Iterate through the string using the right pointer
	for right < len(s) {
		// First, build the initial window of size k
		if right < k {
			if isVowel(s[right]) {
				count++
				maxVowels += 1
			}
			right++
			continue
		}

		// Slide the window:
		// Remove the leftmost character from the window
		if isVowel(s[left]) {
			count--
		}
		left++

		// Add the new character to the window
		if isVowel(s[right]) {
			count++
		}
		right++

		// Update maxVowels if necessary
		if count > maxVowels {
			maxVowels = count
		}
	}

	// Return the maximum number of vowels found in any substring of length k
	return maxVowels
}
