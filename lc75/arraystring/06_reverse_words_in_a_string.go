package arraystring

import "strings"

// reverseWords reverses the order of words in a given string.
//
// Parameters:
//   - s: A string containing words separated by spaces.
//
// Returns:
//
//	A new string with the words reversed in order, with leading/trailing spaces removed
//	and multiple spaces between words reduced to a single space.
//
// The function splits the input string into tokens, reverses their order,
// and then joins them back together, ensuring proper spacing between words.
func reverseWords(s string) string {
	tokens := strings.Split(s, " ")

	out := ""
	for i := len(tokens) - 1; i >= 0; i-- {
		if tokens[i] == "" {
			continue
		}

		if len(out) > 0 {
			out += " "
		}
		out += tokens[i]
	}
	return out
}
