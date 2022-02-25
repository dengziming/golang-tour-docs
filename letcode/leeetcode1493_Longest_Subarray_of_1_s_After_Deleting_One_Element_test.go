package letcode

import "testing"

func TestLongestSubarray(t *testing.T) {
	if longestSubarray([]int{0,0,1,1}) != 2 {
		t.Error("测试失败")
	}
	if longestSubarray([]int{1,1,0,0}) != 2 {
		t.Error("测试失败")
	}
}
