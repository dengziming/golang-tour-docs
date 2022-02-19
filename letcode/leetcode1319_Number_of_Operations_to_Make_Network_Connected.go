package letcode

func makeConnected(n int, connections [][]int) int {
	headquarter := make([]int, n)
	for i := range headquarter {
		headquarter[i] = i
	}

	var find func(x int) int
	// 返回当前是否已经 union
	var union func(x int, y int)

	find = func(x int) int {
		if headquarter[x] != x {
			headquarter[x] = find(headquarter[x])
		}
		return headquarter[x]
	}

	union = func(x int, y int) {
		headquarter[find(x)] = find(y)
	}

	redundant := 0
	for _, conn := range connections {
		if find(conn[0]) == find(conn[1]) {
			redundant++
		} else {
			union(conn[0], conn[1])
		}
	}

	disconnected := 0
	for i, elem := range headquarter {
		if i == elem {
			disconnected ++
		}
	}

	if redundant >= disconnected - 1 {
		return disconnected - 1
	} else {
		return -1
	}
}
