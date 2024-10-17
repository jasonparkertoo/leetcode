package arraystring

// canPlaceFlowers determines if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
//
// Parameters:
//   - flowerbed: A slice of integers representing the flowerbed, where 0 means an empty plot and 1 means a planted flower.
//   - n: An integer representing the number of new flowers to be planted.
//
// Returns:
//
//	A boolean value: true if it's possible to plant n new flowers without violating the rule, false otherwise.
//
// The function iterates through the flowerbed, checking each empty plot to see if it's valid for planting.
// A plot is valid if it's empty and has no adjacent flowers (including the edges of the flowerbed).
// The function modifies the flowerbed in-place, planting flowers as it goes, and returns true if all n flowers can be planted.
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			left := (i == 0) || (flowerbed[i-1] == 0)
			right := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			if left && right {
				flowerbed[i] = 1
				n--
			}
		}

		if n <= 0 {
			return true
		}
	}

	return false
}
