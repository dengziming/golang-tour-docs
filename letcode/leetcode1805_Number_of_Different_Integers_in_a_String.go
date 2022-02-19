package letcode

func numDifferentIntegers(word string) int {
	left := 0
	right := 0
	counts := make(map[string]int)

	for left < len(word) {
		// 不是数字
		for left < len(word) && (word[left] < '0' || word[left] > '9') {
			left++
		}
		right = left

		// 不是
		for right < len(word) && word[right] >= '0' && word[right] <= '9' {
			right++
		}
		if left < right {
			// 去掉头部的 0
			for left < right && word[left] == '0' {
				left ++
			}
			if left >= right {
				// 避免全 0
				counts["0"]++
			} else {
				counts[word[left: right]]++
			}
		}
		left = right + 1
	}
	return len(counts)
}
