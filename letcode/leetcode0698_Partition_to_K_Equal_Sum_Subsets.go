package letcode

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	used := make([]bool, len(nums))
	target := sum / k

	var backtrack func(k int, bucket int, start int) bool
	backtrack = func(k int, bucket int, start int) bool {
		if k == 0 {
			return true
		}
		if bucket == target {
			if backtrack(k-1, 0, 0) {
				return true
			}
			// return backtrack(k-1, 0, 0)
		}

		for i := start; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if bucket+nums[i] > target {
				continue
			}

			used[i] = true
			bucket += nums[i]
			if backtrack(k, bucket, i+1) {
				return true
			}
			used[i] = false
			bucket -= nums[i]
		}

		return false
	}

	return backtrack(k, 0, 0)
}
