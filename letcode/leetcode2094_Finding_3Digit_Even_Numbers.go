package letcode

func findEvenNumbers(digits []int) []int {
	counts := make(map[int]int)

	for _, digit := range digits {
		counts[digit]++
	}

	result := make([]int, 0)
	for i := 100; i < 1000; i+=2 {
		v1 := i / 100
		v2 := i % 100 / 10
		v3 := i % 10

		counts2 := make(map[int]int)
		counts2[v1] ++
		counts2[v2] ++
		counts2[v3] ++

		flag := true
		for i2 := range counts2 {
			if counts2[i2] > counts[i2] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
	}
	return result
}
