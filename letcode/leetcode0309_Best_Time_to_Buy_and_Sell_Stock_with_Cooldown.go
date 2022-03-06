package letcode

func maxProfit3(prices []int) int {
	// 想贪心，没搞定

	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}

	// dp 吧
	n := len(prices)
	// 买了股票
	dpHas := make([]int, n)
	// 前天之前卖了股票
	dpSell := make([]int, n)
	// 昨天卖了股票
	dpCool := make([]int, n)

	dpHas[0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 昨天不持有股票今天买, 或者昨天就持有股票
		dpHas[i] = max(dpHas[i-1], dpSell[i-1]-prices[i])
		// 前天之前卖了股票
		dpSell[i] = max(dpSell[i-1], dpCool[i-1])
		// 昨天就已经卖了，昨天持有今天卖了
		dpCool[i] = dpHas[i-1]+prices[i]
	}

	return max(dpSell[n-1], dpCool[n-1])
}
