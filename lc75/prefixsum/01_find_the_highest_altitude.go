package prefixsum

func largestAltitude(gain []int) int {
	var maxAlt, alt int = 0, 0
	for _, g := range gain {
		alt += g
		if maxAlt < alt {
			maxAlt = alt
		}
	}
	return maxAlt
}
