package letcode

import "testing"

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
// 链接：https://leetcode-cn.com/problems/word-break
func TestWordBreak(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet","code"}) {
		t.Error("测试失败")
	}

	if !wordBreak("applepenapple", []string{"apple","pen"}) {
		t.Error("测试失败")
	}

	if wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}) {
		t.Error("测试失败")
	}

	if !wordBreak("cars", []string{"car","ca","rs"}) {
		t.Error("测试失败")
	}

	if !wordBreak("a", []string{"a"}) {
		t.Error("测试失败")
	}
}
