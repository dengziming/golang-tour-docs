package letcode

func gardenNoAdj(n int, paths [][]int) []int {
	// 先得到每个点的邻居
	// dfs 每个点，如果某个颜色没有用过，就给这个点这个一个颜色。并且给每个邻居添加标识，代表这个颜色已经用过。

	adjs := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		adjs[i] = make([]int, 0)
	}
	for _, path := range paths {
		adjs[path[0]] = append(adjs[path[0]], path[1])
		adjs[path[1]] = append(adjs[path[1]], path[0])
	}

	// i 当前遍历到哪个点，从 1 开始
	var dfs func(i int, target []int) []int

	flag := make([][]int, n+1)
	for i := range flag {
		flag[i] = make([]int, 5) // 代表 1，2，3，4 四种花是否用过
	}

	//
	canUse := func(i int, color int) bool {
		for _, adj := range adjs[i] {
			if flag[adj][color] == 1 {
				return false
			}
		}
		return true
	}

	dfs = func(i int, target []int) []int {
		if i == n + 1 {
			return target
		}
		// 如果所有的邻居都没有用这个色
		for color := 1; color <= 4; color++ {
			if canUse(i, color) {
				flag[i][color] = 1
				tmp := dfs(i + 1, append(target, color))
				flag[i][color] = 0
				if tmp != nil {
					return tmp
				}
			}
		}
		return nil
	}

	return dfs(1, []int{})
}
