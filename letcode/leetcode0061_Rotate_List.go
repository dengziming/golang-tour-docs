package letcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k==0 {
		return head
	}
	array := make([]*ListNode,0)
	for head != nil {
		array = append(array, head)
		head = head.Next
	}
	k = k % len(array)
	if k == 0 {
		return array[0]
	}

	array[len(array)-k-1].Next = nil
	array[len(array)-1].Next = array[0]
	return array[len(array)-k]
}
