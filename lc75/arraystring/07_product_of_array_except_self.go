package arraystring

// productExceptSelf calculates the product of all elements in the input array except the element at each index.
//
// Parameters:
//   - nums: A slice of integers representing the input array.
//
// Returns:
//
//	A slice of integers where each element at index i is the product of all elements in nums except nums[i].
//
// The function uses two passes through the array:
// 1. Left-to-right pass to calculate the product of all elements to the left of each index.
// 2. Right-to-left pass to multiply the left products with the product of all elements to the right of each index.
// This approach avoids using division and achieves O(n) time complexity with O(1) extra space (excluding the output array).
func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))

	left := 1
	for j := range nums {
		out[j] = left * 1
		left *= nums[j]
	}

	right := 1
	for k := len(nums) - 1; k >= 0; k-- {
		out[k] *= right
		right *= nums[k]
	}

	return out
}