package letcode

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	if simplifyPath("/home/") != "/home" {
		t.Error("测试失败")
	}
}