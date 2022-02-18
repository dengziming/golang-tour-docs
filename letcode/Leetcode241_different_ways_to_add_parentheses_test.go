package letcode

import (
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {

	if len(diffWaysToCompute("2-1-1")) != 2 {
		t.Error("测试失败")
	}

	if len(diffWaysToCompute("2*3-4*5")) != 5 {
		t.Error("测试失败")
	}

	if len(diffWaysToCompute("11")) != 1 {
		t.Error("测试失败")
	}

}
