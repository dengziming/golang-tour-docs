package letcode

import (
	"testing"
)

func TestMinFallingPathSum(t *testing.T) {
	if minFallingPathSum([][]int{{1,2,3}, {4,5,6}, {7,8,9}}) != 13 {
		t.Error("测试失败")
	}

}
