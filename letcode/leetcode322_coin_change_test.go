package letcode

import "testing"

func TestCoinChange(t *testing.T) {

	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Error("测试失败")
	}

	if coinChange([]int{2}, 3) != -1 {
		t.Error("测试失败")
	}

	if coinChange([]int{1}, 0) != 0 {
		t.Error("测试失败")
	}

	if coinChange([]int{270,373,487,5,62}, 8121) != 21 {
		t.Error("测试失败")
	}

}
