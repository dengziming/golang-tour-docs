package letcode

import (
	"fmt"
	"testing"
)

func TestNearestPalindromic(t *testing.T) {
	if fmt.Sprint(1) != "1" {
		println(fmt.Sprint(1))
		t.Error("测试失败")
	}
	if nearestPalindromic("123") != "121" {
		t.Error("测试失败")
	}
	if nearestPalindromic("1") != "0" {
		t.Error("测试失败")
	}
	if nearestPalindromic("2") != "1" {
		t.Error("测试失败")
	}
	if nearestPalindromic("11") != "9" {
		t.Error("测试失败")
	}
}
