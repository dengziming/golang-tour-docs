package letcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	maxPathSum1(root, &result)
	return result
}

func maxPathSum1(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := maxPathSum1(root.Left, result)
	right := maxPathSum1(root.Right, result)
	var max int
	var path int
	if left <= 0 && right <= 0 {
		max = root.Val
		path = root.Val
	} else if left <= 0 {
		max = root.Val+ right
		path = root.Val+ right
	} else if right <= 0 {
		max =  root.Val + left
		path =  root.Val + left
	} else {
		max =  root.Val + left + right
		if left < right {
			path = root.Val + right
		} else {
			path = root.Val + left
		}
	}
	if max > *result {
		*result = max
	}
	return path
}
