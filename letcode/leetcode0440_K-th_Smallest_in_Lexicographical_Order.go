package letcode

func findKthNumber(n, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	getSteps := func(cur, n int) (steps int) {
		first, last := cur, cur
		for first <= n {
			steps += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return
	}

	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

// 写错了，写的是所有数按照字典排序后，小于 k 有多少个数
/*func findKthNumber(n, k int) int {
	// [1, n] 之间排序
	// 例如 k = 789, 所有首位小于等于 7 的数

	str := strconv.Itoa(k)
	length := len(str)
	if length == 1 {
		return k
	}

	result := 0
	topV := 1
	for i := length; i > 0; i-- {
		cnt ,_ := strconv.Atoi(str[i-1:i])

		// 最高位少一个数
		if i == 1 {
			cnt -= 1
		}
		result += topV * cnt
		topV = topV*10+1
	}
	return result
}*/