package letcode

func numRabbits(answers []int) int {
	counts := make(map[int]int)

	for _, num := range answers {
		counts[num+1] += 1
	}

	result := 0
	for num, count := range counts {
		group := count / num
		result += group * num

		if count%num > 0 {
			result += num
		}
	}

	return result
}
