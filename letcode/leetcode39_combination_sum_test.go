package letcode

import "testing"

func TestCombinationSum(t *testing.T) {
	if len(combinationSum([]int{2, 3, 6, 7}, 7)) != 2 {
		printlnArray(combinationSum([]int{2, 3, 6, 7}, 7))
		t.Error("测试失败")
	}

	printlnArray(combinationSum([]int{2,3,5}, 8))
	if len(combinationSum([]int{2,3,5}, 8)) != 3 {
		t.Error("测试失败")
	}
}
