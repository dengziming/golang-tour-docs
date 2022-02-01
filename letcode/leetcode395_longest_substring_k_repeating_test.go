package letcode

import "testing"

func TestLongestSubstring(t *testing.T) {
	if longestSubstring("aaabb", 3) != 3{
		t.Error("测试失败")
	}

}
