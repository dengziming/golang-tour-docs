package letcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	// 每个数字的 父节点
	parents := make(map[int]*TreeNode)
	// 每个数字的节点
	nodes := make(map[int]*TreeNode)

	// 构造 tree
	for _, des := range descriptions {
		parent, child, left := des[0], des[1], des[2]

		if _, ok := nodes[parent]; !ok {
			nodes[parent] = &TreeNode{parent, nil, nil}
		}
		if _, ok := nodes[child]; !ok {
			nodes[child] = &TreeNode{child, nil, nil}
		}

		if left == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}

		parents[child] = nodes[parent]
	}

	for child := range nodes {
		if parents[child] == nil {
			return nodes[child]
		}
	}
	return nil
}
