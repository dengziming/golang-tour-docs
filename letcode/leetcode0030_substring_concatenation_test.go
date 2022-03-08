package letcode

import "testing"

func TestFindSubstring(t *testing.T) {
	v1 := findSubstring("barfoothefoobarman", []string{"foo","bar"})
	if len(v1) != 2 || v1[0] != 0 || v1[1] != 9 {
		t.Error("测试失败")
	}


	v2 := findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"})
	if len(v2) != 0 {
		t.Error("测试失败")
	}

	v3 := findSubstring("barfoofoobarthefoobarman", []string{"bar","foo","the"})
	if len(v3) != 3 && v3[0] != 6 || v3[1] != 9 || v3[2] != 12 {
		t.Error("测试失败")
	}

	v4 := findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","good"})
	println(v4)
	if len(v4) != 1 && v4[0] != 8 {
		t.Error("测试失败")
	}

}
