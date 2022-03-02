package letcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	if removeDuplicates("abcd", 2) != "abcd" {
		t.Error("测试失败")
	}
}
