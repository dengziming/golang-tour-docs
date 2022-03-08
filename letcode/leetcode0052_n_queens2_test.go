package letcode

import (
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	if totalNQueens(4) != 2 {
		t.Error("测试失败")
	}
}
