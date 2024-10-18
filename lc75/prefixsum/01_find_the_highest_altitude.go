package prefixsum

// largestAltitude calculates the highest altitude reached based on a series of altitude gains.
// It takes a slice of integers 'gain' representing the net altitude change between two points.
// The function returns the maximum altitude reached as an integer.
//
// The initial altitude is assumed to be 0, and each element in 'gain' represents the
// change in altitude from the previous point. The function keeps track of the current
// altitude and updates the maximum altitude whenever a new peak is reached.
//
// Parameters:
//   - gain: A slice of integers representing altitude changes.
//
// Returns:
//   - The highest altitude reached as an integer.
func largestAltitude(gain []int) int {
	// Initialize variables to track the maximum altitude (maxAlt) and current altitude (alt)
	var maxAlt, alt int = 0, 0

	// Iterate through each altitude gain in the input slice
	for _, g := range gain {
		// Update the current altitude by adding the gain
		alt += g

		// If the current altitude is higher than the maximum recorded altitude,
		// update the maximum altitude
		if maxAlt < alt {
			maxAlt = alt
		}
	}

	// Return the highest altitude reached
	return maxAlt
}
