package letcode

import "testing"

func TestKthPalindrome(t *testing.T) {
	if len(kthPalindrome([]int{492605370,206710368,19,985427531,55,13,979243001,831564215,83}, 15)) != 9 {
		t.Error("测试失败")
	}
	if len(kthPalindrome([]int{2,4,6}, 4)) != 3 {
		t.Error("测试失败")
	}
	if len(kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3)) != 6 {
		t.Error("测试失败")
	}
}
