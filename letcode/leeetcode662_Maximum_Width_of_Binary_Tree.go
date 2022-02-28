package letcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	root.Val = 0
	queue := []*TreeNode{root}

	result := 0
	for len(queue) != 0 {
		var tmp []*TreeNode
		min, max := 0, 0
		for i, node := range queue {
			if node.Left != nil {
				node.Left.Val = node.Val * 2
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Val * 2 + 1
				tmp = append(tmp, node.Right)
			}
			if i == 0 {
				min = node.Val
			}
			if i == len(queue)-1 {
				max = node.Val
			}
		}
		if max - min+1 > result {
			result = max - min+1
		}
		queue = tmp
	}
	return result
}
