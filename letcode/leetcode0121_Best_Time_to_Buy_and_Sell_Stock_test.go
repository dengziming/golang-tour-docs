package letcode

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	if maxProfit([]int{7,1,5,3,6,4}) != 5 {
		t.Error("测试失败")
	}

	if maxProfit([]int{7,6,4,3,1}) != 0 {
		t.Error("测试失败")
	}
}
