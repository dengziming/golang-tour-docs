package letcode

import (
	"math"
)

// https://leetcode-cn.com/problems/number-of-ways-to-reconstruct-a-tree/
func checkWays(pairs [][]int) int {
	// 记录每个节点所有的 祖先和子孙节点(下面统称为 degree 节点)
	adj := make(map[int]map[int]bool)

	for _, pair := range pairs {
		x, y := pair[0], pair[1]
		if adj[x] == nil {
			adj[x] = map[int]bool{}
		}
		adj[x][y] = true
		if adj[y] == nil {
			adj[y] = map[int]bool{}
		}
		adj[y][x] = true
	}
	n := len(adj)

	// root 一定是最多的，所以看看 pair 节点最多的有几个。分别设置为 root 节点，然后去掉它，转化为一个子问题
	// 如果用 treeMap 直接排序就好了，哎 😌
	root := -1
	for node, neighbours := range adj {
		if len(neighbours) >= n {
			return 0
		}
		if len(neighbours) == n-1 {
			root = node
			break
		}
	}

	if root == -1 {
		return 0
	}

	ans := 1
	for node, neighbours := range adj {
		if node == root {
			continue
		}

		currDegree := len(neighbours)
		parent := -1
		parentDegree := math.MaxInt32
		// 根据 degree 的大小找到 node 的父节点 parent
		for neighbour := range neighbours {
			if len(adj[neighbour]) < parentDegree && len(adj[neighbour]) >= currDegree {
				parent = neighbour
				parentDegree = len(adj[neighbour])
			}
		}
		if parent == -1 {
			return 0
		}
		// 检测 neighbours 是否为 adj[parent] 的子集
		for neighbour := range neighbours {
			if neighbour != parent && !adj[parent][neighbour] {
				return 0
			}
		}

		if parentDegree == currDegree {
			ans = 2
		}

	}
	return ans
}
