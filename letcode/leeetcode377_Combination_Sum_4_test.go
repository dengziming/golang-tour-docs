package letcode

import "testing"

func TestCombinationSum4(t *testing.T) {
	if combinationSum4([]int{1,2,3}, 4) != 7 {
		t.Error("测试失败")
	}

}
