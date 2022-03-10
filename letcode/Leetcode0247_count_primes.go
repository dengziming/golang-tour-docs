package letcode

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	if n == 3 {
		return 1
	}

	/*var isPrime func(n int) bool
	isPrime = func(n int) bool {
		is := true
		for i := 2; i * i <= n; i++ {
			if n%i == 0 {
				is = false
			}
		}
		return is
	}*/

	prime := make([]bool, n)
	for i := 0; i < n; i++ {
		prime[i] = true
	}

	result := 0
	for i := 2; i < n; i++ {
		if prime[i] {
			result++
			for j :=  i; j * i < n; j++ {
				prime[j * i] = false
			}
		}
	}

	return result
}
