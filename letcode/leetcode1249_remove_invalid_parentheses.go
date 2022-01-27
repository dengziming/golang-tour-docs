package letcode

import "strings"

// a()b()c(d)ee)ff(
// 整体思路:
//  1. 从左到右遍历，遇到左括号进栈，右括号出栈
//  2. 进栈的是 index，如果出栈失败代表是没有 cp 的右括号，最后栈剩下的是 cp 的左括号。
// 思路2: 直接从左到右，删掉多余的 右括号，然后从右到左删除多余的左括号。太粗暴了忽略。
func minRemoveToMakeValid(s string) string {
	var stack []int
	var remove []int
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		}
		if c == ')' {
			if len(stack) == 0 {
				remove = append(remove, i)
			} else {
				stack = stack[:len(stack) - 1]
			}
		}
	}
	remove = append(remove, stack...)

	var result strings.Builder
	start := 0
	end := -1
	for _, index := range remove {
		end = index
		result.WriteString(s[start: end])
		start = index + 1
	}
	// 最后一段也补上来
	result.WriteString(s[end + 1:])

	return result.String()
}
