package scratch

func slidingWindow(data []int, window int) int {

	// check that data is greater than the window size
	if len(data) < window {
		return -1
	}

	sum, maxSum := 0, -1<<63
	for _, n := range data[:window] {
		sum += n
	}

	if sum > maxSum {
		maxSum = sum
	}

	for i, n := range data[window:] {
		sum -= data[window - window + i]
		sum += n
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}