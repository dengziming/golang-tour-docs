package letcode

import (
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	if len(solveNQueens(4)) != 2 {
		t.Error("测试失败")
	}

}
