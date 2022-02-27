package letcode

import (
	"testing"
)

func TestMinimumDifference(t *testing.T) {
	if minimumDifference([]int{3,1,2}) != -1 {
		t.Error("测试失败")
	}
}
