package letcode

import "testing"


func TestCanJump(t *testing.T) {
	if !canJump([]int{2,3,1,1,4}) {
		t.Error("测试失败")
	}
}