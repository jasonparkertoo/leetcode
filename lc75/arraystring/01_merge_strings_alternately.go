package arraystring

// mergeAlternately combines two strings by alternately taking characters from each string.
//
// Parameters:
//   - word1: The first input string.
//   - word2: The second input string.
//
// Returns:
//
//	A new string that contains characters from word1 and word2 alternately.
//	If one string is longer than the other, the remaining characters are appended at the end.
//
// The function uses two pointers to keep track of the current position in each input string
// and builds the result string by alternating between characters from each input.
func mergeAlternately(word1, word2 string) string {
	w1Pos, w2Pos := 0, 0
	left := true

	merged := make([]byte, len(word1)+len(word2))
	for i := 0; i < (len(word1) + len(word2)); i++ {
		if left {
			merged[i] = word1[w1Pos]
			w1Pos++
			if w2Pos < len(word2) {
				left = false
			}
		} else {
			merged[i] = word2[w2Pos]
			w2Pos++
			if w1Pos < len(word1) {
				left = true
			}
		}
	}

	return string(merged)
}