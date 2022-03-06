package letcode

import (
	"testing"
)

func TestMinMovesToMakePalindrome(t *testing.T) {
	if minMovesToMakePalindrome("roggbxxpbpro") != 10 {
		println(minMovesToMakePalindrome("roggbxxpbpro"))
		t.Error("测试失败")
	}
	if minMovesToMakePalindrome("scpcyxprxxsjyjrww") != 42 {
		println(minMovesToMakePalindrome("scpcyxprxxsjyjrww"))
		t.Error("测试失败")
	}
	if minMovesToMakePalindrome("eqvvhtcsaaqtqesvvqch") != 17 {
		println(minMovesToMakePalindrome("eqvvhtcsaaqtqesvvqch"))
		t.Error("测试失败")
	}
}
