package letcode

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	// 所有的父节点
	father := make([][]int, n)
	for i := 0; i < n; i++ {
		father[i] = make([]int,0)
	}

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		father[to] = append(father[to], from)
	}

	cache := make([][]int, n)

	RemoveDuplicate := func(ints []int) []int {
		if len(ints) == 0 {
			return ints
		}
		result := []int{ints[0]}
		pre := ints[0]
		for i := 1; i < len(ints); i++ {
			if pre != ints[i] {
				result = append(result, ints[i])
				pre = ints[i]
			}

		}
		return result
	}

	var getAns func(node int) []int
	getAns = func(node int) []int {
		// 没有父节点，直接返回空？
		if len(father[node]) == 0 {
			return []int{}
		}

		if cache[node] != nil {
			return cache[node]
		}

		result := append([]int{}, father[node]...)
		for _, fnode := range father[node] {
			result = append(result, getAns(fnode)...)
		}
		sort.Ints(result)
		cache[node] = RemoveDuplicate(result)
		return result
	}
	for i := 0; i < n; i++ {
		getAns(i)
	}
	return cache
}
