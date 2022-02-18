package letcode

import "testing"


func TestCountPrimes(t *testing.T) {
	if countPrimes(10) != 4 {
		t.Error("测试失败")
	}

	if countPrimes(0) != 0 {
		t.Error("测试失败")
	}

	if countPrimes(1) != 0 {
		t.Error("测试失败")
	}
}
