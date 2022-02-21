package letcode

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]byte, 0)
	last := func() byte {
		return stack[len(stack) - 1]
	}
	push := func(b byte) bool {
		if b == '(' || b == '{' || b == '[' {
			stack = append(stack, b)
			return true
		} else {
			if len(stack) == 0 {
				return false
			} else {
				l := last()
				if l == '(' && b == ')' || l == '{' && b == '}' || l == '[' && b == ']'{
					stack = stack[:len(stack) - 1]
					return true
				} else {
					return false
				}
			}
		}
	}

	for i := range s {
		if !push(s[i]) {
			return false
		}
	}
	return len(stack) == 0
}
