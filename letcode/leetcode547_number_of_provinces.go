package letcode

// 并查集
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	headquarter := make([]int, n)
	for i := range headquarter {
		headquarter[i] = i
	}

	var union func(i int, j int)
	var find func(i int) int

	union = func(i int, j int)  {
		headquarter[find(i)] = find(j)
	}

	find = func(i int) int {
		if headquarter[i] == i {
			return i
		}
		head := find(headquarter[i])
		headquarter[i] = head
		return head
	}

	for i := range isConnected {
		// 每一轮下来找到 headquarter = i 的
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				// 合并两个
				union(i, j)
			}
		}
	}

	result := 0
	for i, j := range headquarter {
		if i == j {
			result++
		}
	}
	return result
}
