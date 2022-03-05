package letcode

import (
	"testing"
)


func TestAtMostNGivenDigitSet(t *testing.T) {
	if atMostNGivenDigitSet([]string{"1","3","5","7"}, 100) != 20 {
		t.Error("测试失败")
	}
	if atMostNGivenDigitSet([]string{"1","4","9"}, 1000000000) != 29523 {
		t.Error("测试失败")
	}
	if atMostNGivenDigitSet([]string{"3","4","8"}, 4) != 2 {
		t.Error("测试失败")
	}
	if atMostNGivenDigitSet([]string{"1","2","3","6","7","8"}, 211) != 79 {
		t.Error("测试失败")
	}

	if atMostNGivenDigitSet([]string{"1","5","7","8"}, 10212) != 340 {
		t.Error("测试失败")
	}
}
