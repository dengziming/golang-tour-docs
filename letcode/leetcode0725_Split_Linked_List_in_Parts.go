package letcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	n := len(arr)
	cnt := n / k
	plus := n % k
	result := make([]*ListNode, k)

	for i := 0; i < k; i++ {
		if i < plus {
			if i*(cnt+1) != 0 {
				arr[i*(cnt+1)-1].Next = nil
			}
			result[i] = arr[i*(cnt+1)]
		} else {
			if i*cnt+plus < n {
				if i*cnt+plus != 0 {
					arr[i*cnt+plus-1].Next = nil
				}
				result[i] = arr[i*cnt+plus]
			}
		}
	}
	return result
}
