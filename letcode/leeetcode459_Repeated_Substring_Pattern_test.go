package letcode

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {

}

func TestKmp(t *testing.T) {
	if kmp("abcabc", "abcabdc") {
		t.Error("测试失败")
	}
	if kmp("ababc", "ababcabaa") {
		t.Error("测试失败")
	}
}
