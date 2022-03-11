package letcode

func removeStones(stones [][]int) int {
	n := len(stones)
	headquarter := make([]int, n)
	for i := 0; i < n; i++ {
		headquarter[i] = i
	}

	var union func(x int, y int)
	var find func(x int) int

	union = func(x int, y int) {
		headquarter[find(x)] = find(y)
	}

	find = func(x int) int {
		if x == headquarter[x] {
			return x
		}
		headquarter[x] = find(headquarter[x])
		return headquarter[x]
	}

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				union(i, j)
			}
		}
	}

	result := 0
	for i, j := range headquarter {
		if i != j {
			result++
		}
	}
	return result
}
