package letcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// 每次走一步
	p1 := head
	// 往前两次
	p2 := head

	for true {
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			break
		}
	}
	return p1
}