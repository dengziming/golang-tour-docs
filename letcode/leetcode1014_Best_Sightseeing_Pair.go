package letcode


func maxScoreSightseeingPair(values []int) int {
	n := len(values)

	max := func(a int, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	currentV := values[0]
	maxV := values[0]
	for i := 1; i < n; i++ {
		currentV = max(values[i] + values[i - 1] - 1, currentV + values[i] - values[i - 1] - 1)
		maxV = max(maxV, currentV)
	}

	return maxV
}
