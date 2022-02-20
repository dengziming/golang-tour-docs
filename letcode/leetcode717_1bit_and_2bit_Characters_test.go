package letcode

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	if !isOneBitCharacter([]int{1,0,0}) {
		t.Error("测试失败")
	}
	if isOneBitCharacter([]int{1,1,1,0}) {
		t.Error("测试失败")
	}
}
