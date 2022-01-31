package letcode

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Error("测试失败")
	}
	if climbStairs(3) != 3 {
		t.Error("测试失败")
	}
}
