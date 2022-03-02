package letcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head
	idx := 0
	for idx < k-1 {
		p1 = p1.Next
		idx++
	}
	// 此时 p1 指向第 k 个节点
	node1 := p1
	value1 := p1.Val
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	value2 := p2.Val

	node1.Val = value2
	p2.Val = value1
	return head
}
