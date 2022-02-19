package letcode

import "testing"

func TestMiddleNode(t *testing.T) {
	if middleNode(&ListNode{1, nil}).Val != 1 {
		t.Error("测试失败")
	}
}
