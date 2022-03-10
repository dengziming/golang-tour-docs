package letcode

import "fmt"

func getHint(secret string, guess string) string {
	// 记录 cows 个数
	count1 := make([]int, 10)
	count2 := make([]int, 10)

	a := 0

	n := len(secret)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			count1[secret[i] - '0']++
			count2[guess[i] - '0']++
		}
	}
	b := 0
	for i := 0; i < 10; i++ {
		if count1[i] > count2[i] {
			b += count2[i]
		} else {
			b += count1[i]
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
