package letcode

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		 return 0
	}
	result := make([]int, amount + 1)

	result[0] = 0

	for i := 1; i <= amount; i++ {
		result[i] = -1
		for _, coin := range coins {
			j := i - coin

			if j >= 0 && result[j] != -1 && (result[j] + 1 < result[i] || result[i] == -1) {
				result[i] = result[j] + 1
			}
		}

	}
	return result[amount]
}
