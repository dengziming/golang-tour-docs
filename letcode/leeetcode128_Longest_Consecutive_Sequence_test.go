package letcode

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	if longestConsecutive([]int{100,4,200,1,3,2}) != 4 {
		t.Error("测试失败")
	}

	if longestConsecutive([]int{0}) != 1 {
		t.Error("测试失败")
	}

	if longestConsecutive([]int{1,2,3,4}) != 4 {
		t.Error("测试失败")
	}

}
