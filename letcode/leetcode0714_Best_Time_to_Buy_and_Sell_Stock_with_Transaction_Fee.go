package letcode

func maxProfit2(prices []int, fee int) int {
	// 第一次买的时候贵 2 元
	// 卖完立即买是原价

	profit := 0
	buy := prices[0] + 2

	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			// 卖完立即买
			profit += prices[i] - buy
			buy = prices[i]
		} else if prices[i] + fee < buy {
			// 这个价格更实惠
			buy = prices[i] + fee
		}
	}
	return profit
}
