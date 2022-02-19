package letcode


func eventualSafeNodes(graph [][]int) []int {
	var dfs func(index int, stopIndex *map[int]int) bool
	// 记忆化搜索
	cache := make([]int, len(graph))
	for i := range cache {
		cache[i] = -1 // 代表未计算
	}

	set := func(k int, v int) {
		cache[k] = v
	}

	dfs = func(index int, stopIndex *map[int]int) bool {
		if cache[index] == 0 {
			return false
		}
		if cache[index] == 1 {
			return true
		}

		var result = true
		if len(graph[index]) == 0 {
			result =  true
		} else if (*stopIndex)[index] == 1 {
			// 说明从一个节点又回到了这个节点，那一路上的节点都不可以
			for i, v := range *stopIndex {
				if v == 1 {
					set(i, 0) // 代表不可以
				}
			}
			result = false
		} else {
			for _, out := range graph[index] {
				(*stopIndex)[index] = 1
				if !dfs(out, stopIndex) {
					result = false
					break
				}
				(*stopIndex)[index] = 0
			}
		}

		// 设置 cache
		if result {
			cache[index] = 1
		} else {
			cache[index] = 0
		}
		return result
	}

	result := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		stopIndex := make(map[int]int)
		if dfs(i, &stopIndex) {
			result = append(result, i)
		}
	}
	return result
}
