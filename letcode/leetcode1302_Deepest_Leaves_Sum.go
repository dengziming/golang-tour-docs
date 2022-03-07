package letcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}

	result := 0
	for len(queue) != 0 {
		tmp := queue
		queue = nil
		v := 0
		for _, node := range tmp {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			v += node.Val
		}
		result = v
	}
	return result
}
