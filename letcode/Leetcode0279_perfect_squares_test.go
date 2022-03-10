package letcode

import (
	"testing"
)


func TestNumSquares(t *testing.T) {

	if numSquares(12) != 3 {
		t.Error("测试失败")
	}

}
