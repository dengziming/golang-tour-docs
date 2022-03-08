package letcode

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	candles := make([]int, 0)
	sum := make([]int, n+1)
	for i := range s {
		if s[i] == '|' {
			candles = append(candles, i)
		}
		sum[i+1] = sum[i]
		if s[i] == '*' {
			sum[i+1]+=1
		}
	}

	// 找到大于等于 i 的第一个蜡烛
	binary1 := func(i int) int {
		left, right := 0, len(candles)-1
		for left < right {
			mid := (left + right)/2
			if candles[mid] >= i {
				right = mid
			} else {
				left = mid+1
			}
		}
		if candles[left] >= i {
			return left
		}
		return -1
	}

	// 找到小于等于 j 的第一个蜡烛
	binary2 := func(j int) int {

		left, right := 0, len(candles)-1
		for left < right {
			mid := (left + right + 1)/2
			if candles[mid] <= j {
				left = mid
			} else {
				right = mid-1
			}
		}
		if candles[right] <=j {
			return right
		}
		return -1
	}

	get := func(i, j int) int {
		if len(candles) == 0 {
			return 0
		}
		// 找 i j 区间内所有的蜡烛即可
		left, right := binary1(i), binary2(j)
		if left == -1 || right == -1 || left >= right {
			return 0
		}
		return sum[candles[right]+1]-sum[candles[left]]
	}

	result := make([]int, 0)
	for _, query := range queries {
		result = append(result, get(query[0], query[1]))
	}
	return result
}
