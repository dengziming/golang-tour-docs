package letcode

import "testing"

func TestLongestNiceSubstring(t *testing.T) {
	if longestNiceSubstring("YazaAay") != "aAa" {
		t.Error("测试失败")
	}

}
