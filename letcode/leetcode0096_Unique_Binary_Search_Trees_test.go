package letcode

import "testing"

func TestNumTrees(t *testing.T) {
	if numTrees(1) != 1 {
		t.Error("测试失败")
	}

	if numTrees(3) != 5 {
		t.Error("测试失败")
	}
}
