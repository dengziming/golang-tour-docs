package letcode

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	v1 := uniquePaths(1, 2)
	if v1 != 1 {
		t.Error("测试失败")
	}

	v2 := uniquePaths(2, 2)
	if v2 != 2 {
		t.Error("测试失败")
	}

}
