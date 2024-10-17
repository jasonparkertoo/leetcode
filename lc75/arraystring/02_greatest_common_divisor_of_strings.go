package arraystring

// gcdOfStrings finds the largest string that divides both input strings.
//
// Parameters:
//   - str1: The first input string.
//   - str2: The second input string.
//
// Returns:
//
//	A string representing the greatest common divisor of str1 and str2.
//	If no such string exists, it returns an empty string.
//
// The function first checks if the concatenation of str1 and str2 is equal to
// the concatenation of str2 and str1. If not, it means there's no common divisor.
// Then, it finds the GCD of the lengths of str1 and str2 using the Euclidean algorithm.
// The substring of str1 with length equal to this GCD is the result.
func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcd := len(str1)
	if len(str2) > gcd {
		gcd = len(str2)
	}
	for gcd > 0 {
		if len(str1)%gcd == 0 && len(str2)%gcd == 0 {
			break
		}
		gcd--
	}
	return str1[:gcd]
}