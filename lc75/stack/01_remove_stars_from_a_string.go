package stack

// removeStars removes characters that precede asterisks (*) in a given string.
//
// Given a string s that contains characters and asterisks (*), for each asterisk,
// the function removes the closest non-asterisk character to its left.
//
// If there is no non-asterisk character to the left of an asterisk, the function
// skips that asterisk.
//
// Parameters:
//   - s: input string containing characters and asterisks
//
// Returns:
//   - string: the final string after processing all asterisks
//
// Example:
//   removeStars("leet**cod*e") returns "lecoe"
//   removeStars("erase*****") returns ""
// 
// Time Complexity: O(n), where n is the length of input string s
//   - The function iterates through each character in the string exactly once
//   - Stack operations (append and removal) are O(1)
//
// Space Complexity: O(n)
//   - In the worst case (no asterisks), the stack will store all n characters
//   - In best case (all asterisks), space used will be O(1)
func removeStars(s string) string {	
	// Initialize a stack to store characters
	var stack []rune
	
	// Iterate through each character in the input string
	for _, char := range s {
		// If current char is '*' and stack is not empty, remove last char
		if char == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		// If current char is not '*', add it to stack
		if char != '*' {
			stack = append(stack, char)
		}
	}
	
	// Convert final stack back to string and return
	return string(stack)
}