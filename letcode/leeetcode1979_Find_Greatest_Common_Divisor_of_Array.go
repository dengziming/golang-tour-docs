package letcode

func findGCD(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxV, minV := nums[0], nums[0]
	for _, num := range nums {
		if num > maxV {
			maxV = num
		}
		if num < minV {
			minV = num
		}
	}

	gcd := func(a,b int) int {
		for b != 0 {
			a,b = b, a % b
		}
		return a
	}
	return gcd(maxV, minV)
}
