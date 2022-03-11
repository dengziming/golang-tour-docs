package letcode

import (
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	if complexNumberMultiply("1+1i", "1+1i") != "0+2i" {
		t.Error("测试失败")
	}
}
