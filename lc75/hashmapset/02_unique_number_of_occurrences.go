package hashmapset

func uniqueOccurrences(arr []int) bool {
	table := map[int]int{}
	for _, n := range arr {
		table[n] += 1
	}

	set := map[int]bool{}
	for _, count := range table {
		if set[count] {
			return false
		}
		set[count] = true
	}
	return true
}
