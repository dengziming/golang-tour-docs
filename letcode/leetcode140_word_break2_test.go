package letcode

import "testing"

func TestWordBreak2(t *testing.T) {

	if len(wordBreak2("leetcode", []string{"leet","code"})) != 1 {
		t.Error("测试失败")
	}

	if len(wordBreak2("pineapplepenapple", []string{"apple","pen","applepen","pine","pineapple"})) != 3 {
		t.Error("测试失败")
	}
	if len(wordBreak2("catsandog", []string{"cats","dog","sand","and","cat"})) != 0 {
		t.Error("测试失败")
	}
	if len(wordBreak2("aaaaaaa", []string{"aaaa","aa","a"})) != 31 {
		t.Error("测试失败")
	}
}
