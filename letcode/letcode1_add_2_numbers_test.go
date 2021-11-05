package letcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	printlnListNode(AddTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
	println()
	printlnListNode(AddTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{6, nil}}}))
}
