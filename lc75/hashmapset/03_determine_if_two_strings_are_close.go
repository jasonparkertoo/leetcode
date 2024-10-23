package hashmapset

import "sort"

// closeStrings determines if two strings are considered "close".
// Two strings are considered close if:
// 1. They have the same length.
// 2. They contain the same set of characters.
// 3. The frequency of characters can be rearranged to match each other.
//
// Parameters:
//   - word1: The first string to compare.
//   - word2: The second string to compare.
//
// Returns:
//   - true if the strings are close, false otherwise.
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	for _, ch := range word1 {
		freq1[ch]++
	}
	for _, ch := range word2 {
		freq2[ch]++
	}

	if len(freq1) != len(freq2) {
		return false
	}

	counts1 := make([]int, 0)
	counts2 := make([]int, 0)

	for ch := range freq1 {
		if _, exists := freq2[ch]; !exists {
			return false
		}
		counts1 = append(counts1, freq1[ch])
		counts2 = append(counts2, freq2[ch])
	}

	sort.Ints(counts1)
	sort.Ints(counts2)

	for i := range counts1 {
		if counts1[i] != counts2[i] {
			return false
		}
	}

	return true
}
