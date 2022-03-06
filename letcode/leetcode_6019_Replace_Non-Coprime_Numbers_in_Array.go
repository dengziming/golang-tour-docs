package letcode


func replaceNonCoprimes(nums []int) []int {
	getN := func(a, b int) int {
		if b>a {
			a, b = b, a
		}
		for b != 0 {
			a, b = b, a % b
		}
		return a
	}

	stack := make([]int, 0)
	var push func(num int)
	push = func(num int) {
		if len(stack) == 0 {
			stack = append(stack, num)
			return
		}
		tail := stack[len(stack)-1]
		gcd := getN(num, tail)
		if gcd > 1 {
			stack = stack[:len(stack)-1]
			push(num*tail/gcd)
		} else {
			stack = append(stack, num)
		}
	}

	for _, num := range nums {
		push(num)
	}

	return stack
}
