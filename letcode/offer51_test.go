package letcode

import "testing"

func TestReversePairs(t *testing.T) {
	if reversePairs([]int{7,5,6,4}) != 5 {
		t.Error("测试失败")
	}

}