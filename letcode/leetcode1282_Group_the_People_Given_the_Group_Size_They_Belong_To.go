package letcode

func groupThePeople(groupSizes []int) [][]int {

	idsBySize := make(map[int][]int)
	for id, size := range groupSizes {
		if _, ok := idsBySize[size]; !ok {
			idsBySize[size] = make([]int,0)
		}
		idsBySize[size] = append(idsBySize[size], id)
	}

	result := make([][]int, 0)
	for size, people := range idsBySize {
		group := len(people) / size
		for i := 0; i < group; i++ {
			subGroup := make([]int,0)
			for j := 0; j < size; j++ {
				subGroup = append(subGroup, people[size*i+j])
			}
			result = append(result, subGroup)
		}
	}
	return result
}
