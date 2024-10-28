package stack

import (
	"strings"
	"unicode"
)

// decodeString decodes a string using a run-length encoding format.
// The format supports repeating sequences specified by digits followed by characters
// in square brackets, e.g. "3[a]" decodes to "aaa".
//  
// The encoding rules are:
// - Numbers indicate how many times the content in the following brackets should be repeated
// - Letters are added directly to the output string 
// - Paired square brackets [] contain the substring to be repeated
// - Encoding can be nested, e.g. "2[3[a]]" decodes to "aaaaaa"
//
// Parameters:
//   - s: String to decode with valid run-length encoding format
//
// Returns:
//   - The decoded string with all repetitions expanded
//
// Example:
//   decodeString("3[a]2[bc]") returns "aaabcbc"
//   decodeString("3[a2[c]]") returns "accaccacc"
func decodeString(s string) string {
	// Initialize stacks for keeping track of multipliers and intermediate strings
	numberStack := make([]int, 0, len(s)/4)
	stringStack := make([]string, 0, len(s)/4)

	// Track the current number being parsed and string being built
	var currentNumber int
	
	var builder strings.Builder
	// Pre-allocate builder capacity to optimize memory usage
	builder.Grow(len(s) * 2)

	// Iterate through each character in the input string
	for i := 0; i < len(s); i++ {
		char := s[i]
		switch {
		case unicode.IsDigit(rune(char)):
			// Accumulate digits to build multi-digit multiplier
			// e.g. "12" becomes 12, not 1+2
			currentNumber = currentNumber*10 + int(char-'0')

		case char == '[':
			// Opening bracket means we start a new group
			// Save current state to stack and reset for new group contents
			numberStack = append(numberStack, currentNumber)
			stringStack = append(stringStack, builder.String())
			currentNumber = 0
			builder.Reset()

		case char == ']':
			// Closing bracket means we've finished a group
			// Pop the multiplier and previous string from stacks
			times := numberStack[len(numberStack)-1]
			numberStack = numberStack[:len(numberStack)-1]

			// Get the string to repeat and restore previous string
			repeatStr := builder.String()
			builder.Reset()
			builder.WriteString(stringStack[len(stringStack)-1])
			stringStack = stringStack[:len(stringStack)-1]

			// Append the repeated string to the current result
			builder.WriteString(strings.Repeat(repeatStr, times))

		default:
			// For regular characters (letters), just append to current string
			builder.WriteByte(char)
		}
	}

	// Return the final decoded string
	return builder.String()
}
