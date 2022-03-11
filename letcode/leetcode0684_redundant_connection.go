package letcode

// 并查集
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	tree := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		tree[i] = i
	}

	var find func(int) int
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		tree[x] = y
		return true
	}

	find = func(x int) int {
		if tree[x] == x {
			return x
		}
		tree[x] = find(tree[x])
		return tree[x]
	}

	for _, tuple := range edges {
		if !union(tuple[0], tuple[1]) {
			return tuple
		}
	}
	return nil
}
