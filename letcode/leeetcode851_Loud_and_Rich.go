package letcode

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	// 记录比 i 更有钱的人
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, tp := range richer {
		graph[tp[1]] = append(graph[tp[1]], tp[0])
	}

	for i := 0; i < n; i++ {

	}

	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i] = -1
	}

	// 找到比 i 有钱的最安静的人
	var dfs func(i int) int

	dfs = func(i int) int {
		if cache[i] != -1 {
			return cache[i]
		}
		// 比 i 有钱
		minV := i
		for _, child := range graph[i] {
			if quiet[dfs(child)] < quiet[minV] {
				minV = dfs(child)
			}
		}
		cache[i] = minV
		return cache[i]
	}

	for i := 0; i < n; i++ {
		cache[i] = dfs(i)
	}
	return cache
}
