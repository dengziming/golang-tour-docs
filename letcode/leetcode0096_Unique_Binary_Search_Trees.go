package letcode

func numTrees(n int) int {
	var katelan func(n int) int
	cache := make(map[int]int)

	katelan = func(n int) int {
		if n == 0 {
			return 1
		}
		if cache[n] != 0 {
			return cache[n]
		}
		count := 0
		for i := 1; i < n + 1; i++ {
			count += katelan(i - 1) * katelan(n - i)
		}
		cache[n] = count
		return count
	}

	return katelan(n)
}
