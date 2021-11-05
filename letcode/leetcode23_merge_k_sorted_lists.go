package letcode


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return merge2Lists(lists[0], mergeKLists(lists[1:]))
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge2Lists(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge2Lists(l2.Next, l1)
		return l2
	}
}
