package letcode

func wateringPlants(plants []int, capacity int) int {

	left := capacity
	path := 0
	//
	for i := range plants {
		// 当前在 i - 1 位置，判断剩余的水是否能搞定 i
		if left >= plants[i] {
			path += 1
			left -= plants[i]
		} else {
			left = capacity - plants[i]
			path += 2*i+1
		}
	}
	return path
}
