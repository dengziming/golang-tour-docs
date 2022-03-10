package letcode

import "testing"

func TestLengthOfLIS(t *testing.T) {
	if lengthOfLIS([]int{1, 2, 5}) != 3 {
		t.Error("测试失败")
	}
}
