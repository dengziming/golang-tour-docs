package letcode

func maxProfit(prices []int) int {
	n := len(prices)

	max := func(a int, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	min := func(a int, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < n; i++ {
		// 第 i 天卖掉的最大受益
		minPrice = min(minPrice, prices[i])
		maxProfit = max(maxProfit, prices[i] - minPrice)

	}
	return maxProfit
}
