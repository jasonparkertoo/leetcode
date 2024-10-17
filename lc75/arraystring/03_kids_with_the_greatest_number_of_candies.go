package arraystring

// kidsWithCandies determines if each kid would have the greatest number of candies after receiving extra candies.
//
// Parameters:
//   - candies: A slice of integers representing the number of candies each kid has.
//   - extraCandies: An integer representing the number of extra candies to be given.
//
// Returns:
//
//	A slice of boolean values where true indicates that the kid at that index would have
//	the greatest number of candies after receiving the extra candies.
//
// The function first finds the maximum number of candies any kid currently has.
// Then, it checks for each kid if their current candies plus the extra candies
// would be greater than or equal to this maximum. The result for each kid is stored
// in a boolean slice which is returned.
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxCandies int
	for _, n := range candies {
		if n > maxCandies {
			maxCandies = n
		}
	}
	out := make([]bool, len(candies))
	for i, c := range candies {
		if (c + extraCandies) >= maxCandies {
			out[i] = true
		}
	}
	return out
}
