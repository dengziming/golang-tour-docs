package letcode

import "testing"

func TestPermute(t *testing.T) {
	if len(permute([]int{1})) != 1 {
		t.Error("测试失败")
	}

	if len(permute([]int{1,2})) != 2 {
		t.Error("测试失败")
	}

	if len(permute([]int{1,2,3})) != 6 {
		t.Error("测试失败")
	}
}