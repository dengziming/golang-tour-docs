package letcode

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	if getHint("1807", "7810") != "1A3B" {
		t.Error("测试失败")
	}

	if getHint("1123", "0111") != "1A1B" {
		t.Error("测试失败")
	}

}