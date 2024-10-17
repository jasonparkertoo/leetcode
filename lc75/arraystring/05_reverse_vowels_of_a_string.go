package arraystring

// reverseVowels reverses the order of vowels in a given string while keeping consonants in place.
//
// Parameters:
//   - s: A string containing vowels and consonants.
//
// Returns:
//
//	A new string with the vowels reversed in order and consonants remaining in their original positions.
//
// The function uses a two-pointer approach to find and swap vowels from both ends of the string
// while preserving the position of consonants. It handles both uppercase and lowercase vowels.
func reverseVowels(s string) string {
	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !isVowel(byte(runes[left])) {
			left++
		}
		for left < right && !isVowel(byte(runes[right])) {
			right--
		}
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)
}