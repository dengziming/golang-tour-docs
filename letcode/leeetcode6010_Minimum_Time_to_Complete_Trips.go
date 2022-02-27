package letcode

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)
	left := 0
	right := totalTrips * time[0]

	check := func(mid int) int {
		cnt := 0
		for _, num := range time {
			if num>mid {
				break
			}
			cnt += mid / num
		}
		return cnt
	}

	for left < right {
		mid := left + (right - left)/2
		//
		if check(mid) >= totalTrips {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}
