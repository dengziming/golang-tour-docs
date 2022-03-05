package letcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		 return nil
	}
	queue := []*Node{root}

	for len(queue) != 0 {
		tmp := queue
		queue = nil

		for i := range tmp {
			if tmp[i].Left != nil {
				queue = append(queue, tmp[i].Left)
			}
			if tmp[i].Right != nil {
				queue = append(queue, tmp[i].Right)
			}
			if i != len(tmp) - 1 {
				tmp[i].Next = tmp[i+1]
			}
		}
	}

	return root
}
