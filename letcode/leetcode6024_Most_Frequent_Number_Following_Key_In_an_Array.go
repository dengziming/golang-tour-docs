package letcode


func mostFrequent(nums []int, key int) int {
	n := len(nums)
	counts := make(map[int]int)
	for i := 0; i < n-1; i++ {
		if nums[i] == key {
			counts[nums[i+1]]++
		}
	}
	max := -1
	maxCnt := -1
	for i, cnt := range counts {
		if cnt > maxCnt {
			maxCnt = cnt
			max = i
		}
	}
	return max
}
