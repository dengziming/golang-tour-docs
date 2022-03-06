package letcode

func carPooling(trips [][]int, capacity int) bool {
	counts := make(map[int]int)
	for _, trip := range trips {
		cnt, from, to := trip[0], trip[1],trip[2]
		for i := from; i < to; i++ {
			counts[i] += cnt
		}
	}
	for _, cnt := range counts {
		if cnt > capacity {
			return false
		}
	}
	return true
}
