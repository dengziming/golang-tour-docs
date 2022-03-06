package letcode

func diStringMatch(s string) []int {
	n := len(s)
	result := make([]int, 0)

	var subtract func(i int, start int, end int)

	subtract = func(i int, start int, end int) {
		if start > end {
			return
		}
		if start == end {
			result = append(result, start)
			return
		}
		// 接下来是上升
		if s[i] == 'I' {
			result = append(result, start)
			subtract(i+1, start+1, end)
		} else {
			result = append(result, end)
			subtract(i+1, start, end-1)
		}
	}
	subtract(0, 0, n)
	return result
}
