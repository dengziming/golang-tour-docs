package letcode

func countHighestScoreNodes(parents []int) int {
	// 从每个节点的父节点，可以得到每个节点的子节点，递归得到每个节点的子孙节点数
	n := len(parents)
	if n == 1 {
		return 1
	}
	children := make([][]int, n)
	for i, parent := range parents {
		if parent != -1 {
			children[parent] = append(children[parent], i)
		}
	}

	cache := make([]int, n)
	for i := 0; i < n; i++ {
		cache[i]=-1
	}
	// i 所有的子孙节点
	var dfs func(i int) int

	dfs = func(i int) int {
		if cache[i] != -1 {
			return cache[i]
		}
		result := 0
		for _, child := range children[i] {
			result += 1
			result += dfs(child)
		}
		cache[i] = result
		return result
	}

	dfs(0)

	maxV := 0
	scores := make(map[int]int)
	for i := 0; i < n; i++ {
		parent := cache[0] - cache[i]
		if parent == 0 {
			parent = 1
		}
		left, right := 1, 1
		if len(children[i]) > 0 {
			left = cache[children[i][0]] + 1
		}
		if len(children[i]) > 1 {
			right = cache[children[i][1]] +1
		}
		value := parent * left * right
		scores[value]+=1
		if value > maxV {
			maxV = value

		}
	}
	return scores[maxV]
}
