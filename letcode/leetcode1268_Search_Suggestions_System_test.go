package letcode

import (
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	if len(suggestedProducts([]string{"bags","baggage","banner","box","cloths"}, "bags")) != 4 {
		t.Error("测试失败")
	}
	if len(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse")) != 5 {
		t.Error("测试失败")
	}
}
