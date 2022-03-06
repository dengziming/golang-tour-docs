package letcode

func maximumPopulation(logs [][]int) int {
	// logs[i] 代表 i 出生死亡年份

	offset := 1950
	delta := make([]int, 101)

	for _, log := range logs {
		delta[log[0]-offset]++
		delta[log[1]-offset]--
	}

	mx, res, curr := 0,0,0

	for i := 0; i < 101; i++ {
		curr += delta[i]
		if curr > mx {
			mx = curr
			res = i
		}
	}
	return res+offset
}
