package letcode

func checkValidString(s string) bool {
	// 记录 ( 的 index
	stack1 := make([]int, 0)
	// 记录 * 的 index
	stack2 := make([]int, 0)

	push := func(i int) bool {
		b := s[i]
		if b == '(' {
			stack1 = append(stack1, i)
			return true
		} else if b == '*' {
			stack2 = append(stack2, i)
			return true
		} else {
			if len(stack1) == 0 && len(stack2) == 0 {
				return false
			} else if len(stack1) == 0 {
				stack2 = stack2[:len(stack2) - 1]
				return true
			} else {
				stack1 = stack1[:len(stack1) - 1]
				return true
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if !push(i) {
			return false
		}
	}

	// 比较剩下的 ( 和 * 能不能匹配上
	for len(stack1) > 0 && len(stack2) > 0 {
		if stack2[len(stack2) - 1] < stack1[len(stack1) - 1]  {
			return false
		} else {
			stack1 = stack1[:len(stack1) - 1]
			stack2 = stack2[:len(stack2) - 1]
		}
	}

	return len(stack1) == 0
}
