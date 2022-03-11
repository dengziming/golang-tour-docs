package letcode

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	if numRollsToTarget(30, 30, 500) != 222616187 {
		println(numRollsToTarget(30, 30, 500))
		t.Error("测试失败")
	}
	if numRollsToTarget(2, 6, 7) != 6 {
		t.Error("测试失败")
	}
}
