package letcode

func minAddToMakeValid(s string) int {
	stack := 0
	add := 0
	for i := range s {
		if s[i] == '(' {
			stack++
		} else {
			stack--
		}
		if stack == -1 {
			stack++
			add++
		}
	}
	return add + stack
}
