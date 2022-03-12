package letcode

import (
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	if len(findSubsequences([]int{4,4,3,2,1})) != 1 {
		t.Error("测试失败")
	}
	if len(findSubsequences([]int{4, 6, 7, 7})) != 8 {
		t.Error("测试失败")
	}

}
