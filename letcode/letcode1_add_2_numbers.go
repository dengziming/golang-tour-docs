package letcode


type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)
}

func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	value := l1.Val + l2.Val + carry
	newVal := value % 10
	newCarry := value / 10

	if newCarry == 0 && l1.Next == nil && l2.Next == nil{
		return &ListNode{newVal, nil}
	} else if l1.Next == nil && l2.Next == nil {
		return &ListNode{newVal, &ListNode{1, nil}}
	} else if l1.Next == nil {
		return &ListNode{newVal, addTwoNumbersWithCarry(&ListNode{0, nil}, l2.Next, newCarry)}
	} else if l2.Next == nil {
		return &ListNode{newVal, addTwoNumbersWithCarry(l1.Next, &ListNode{0, nil}, newCarry)}
	}
	return &ListNode{
		newVal,
		addTwoNumbersWithCarry(l1.Next, l2.Next, newCarry),
	}
}
