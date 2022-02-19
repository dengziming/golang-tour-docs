package letcode

import "testing"

func TestGenerateTrees(t *testing.T) {
	if len(generateTrees(1)) != 1 {
		t.Error("测试失败")
	}

	if len(generateTrees(3)) != 5 {
		t.Error("测试失败")
	}
}
