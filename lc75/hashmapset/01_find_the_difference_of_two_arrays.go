package hashmapset

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]struct{}{}
	set2 := map[int]struct{}{}

	answer := [][]int{{},{}}

	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, exists := set1[num]; !exists {
			answer[1] = append(answer[1], num)
		}
		set2[num] = struct{}{}
	}

	for num := range set1 {
		if _, exists := set2[num]; !exists {
			answer[0] = append(answer[0], num)
		}
	}

	return answer
}

