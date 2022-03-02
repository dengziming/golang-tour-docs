package letcode

func removeDuplicates(s string, k int) string {
	// char
	stack1 := make([]uint8,1)
	// count
	stack2 := make([]int,1)

	stack1[0] = '#'
	stack2[0] = 0

	for i := range s {
		if stack1[len(stack1)-1] != s[i] {
			stack1 = append(stack1, s[i])
			stack2 = append(stack2, 1)
		} else {
			stack2[len(stack2)-1]++
			if stack2[len(stack2)-1] == k {
				stack1 = stack1[:len(stack1)-1]
				stack2 = stack2[:len(stack2)-1]
			}
		}

	}
	stack := []byte{}
	for i := 1; i < len(stack1); i++ {
		for j := 0; j < stack2[i]; j++ {
			stack = append(stack, stack1[i])
		}
	}
	return string(stack)
}
