package letcode

import "testing"

func TestMergeKLists(t *testing.T) {
	if mergeKLists([]*ListNode{}) != nil {
		t.Error("测试失败")
	}

	// [1,4,5],[1,3,4],[2,6]
	if mergeKLists([]*ListNode{createListNode([]int{1, 4, 5}), createListNode([]int{1, 3, 4}), createListNode([]int{2, 6})}) != createListNode([]int{1,1,2,3,4,4,5,6}) {
		printlnListNode(mergeKLists([]*ListNode{createListNode([]int{1, 4, 5}), createListNode([]int{1, 3, 4}), createListNode([]int{2, 6})}))
		t.Error("测试失败")
	}
}
