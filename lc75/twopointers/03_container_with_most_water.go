package twopointers

// maxArea calculates the maximum area of water that can be contained between two vertical lines.
//
// Parameters:
//   - height: A slice of integers representing the heights of vertical lines.
//
// Returns:
//
//	An integer representing the maximum area of water that can be contained.
//
// The function uses a two-pointer approach, starting from the leftmost and rightmost lines.
// It calculates the area between these lines and moves the pointer with the shorter height inward.
// This process continues until the pointers meet, keeping track of the maximum area encountered.
func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0

	for left < right {
		width := right - left
		containerHeight := min(height[left], height[right])
		area := width * containerHeight

		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
