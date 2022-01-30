package letcode

import "strings"

// Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences in any order.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
//链接：https://leetcode-cn.com/problems/word-break-ii
func wordBreak2(s string, wordDict []string) []string {
	n := len(s)
	result := make([][][]string, n + 1)

	result[0] = [][]string{ []string{""} }
	for i := 1; i <= n; i ++ {
		result[i] = make([][]string, 0)
		for j := 0; j < i; j++ {
			for _, word := range wordDict {
				if len(word) == (i - j) && word == s[j: i] {
					for _, pre := range result[j] {
						result[i] = append(result[i], append(append([]string{}, pre...), word))
					}

				}
			}
		}
	}
	array := result[n]

	m := make([]string, len(array))
	for i, data := range array {
		m[i] = strings.Join(data[1:], " ")
	}
	return m
}
