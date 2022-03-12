package letcode

import (
	"testing"
)

func TestNthSuperUglyNumber(t *testing.T) {
	if nthSuperUglyNumber(12, []int{2,7,13,19}) != 32 {
		t.Error("测试失败")
	}
}
