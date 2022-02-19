package letcode

import "testing"

func TestNumDifferentIntegers(t *testing.T) {
	if numDifferentIntegers("a123bc34d8ef34") != 3 {
		t.Error("测试失败")
	}

	if numDifferentIntegers("leet1234code234") != 2 {
		t.Error("测试失败")
	}
	if numDifferentIntegers("a1b01c001") != 1 {
		t.Error("测试失败")
	}
	if numDifferentIntegers("9") != 1 {
		t.Error("测试失败")
	}

	if numDifferentIntegers("00") != 1 {
		t.Error("测试失败")
	}

	if numDifferentIntegers("0a0") != 1 {
		t.Error("测试失败")
	}
}
