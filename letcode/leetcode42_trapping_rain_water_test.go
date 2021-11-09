package letcode

import "testing"

func TestTrap(t *testing.T) {

	if trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}) != 6 {
		t.Error("测试失败")
	}
}